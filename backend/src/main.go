package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/implocell/solid-stream/src/ticker"

	"github.com/gorilla/websocket"
)

var SEED int64 = 355015540

var tickers []*ticker.Ticker

var messagesSincePrint int64 = 0

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func onInstrumentUpdate(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Print("Upgrade ERR: ", err)
	}

	defer conn.Close()

	symbol := req.URL.Path[7:len(req.URL.Path)]

	var subject *ticker.Ticker
	for _, ticker := range tickers {
		if ticker.Symbol == symbol {
			subject = ticker
			break
		}
	}

	if subject == nil {
		log.Print("Could not match ticker based on symbol ", symbol)

		return
	}

	log.Print(req.RemoteAddr, " subscribed to ", symbol)

	for {
		currentTime := time.Now().UnixMilli()

		if subject.NextUpdate < currentTime {

			json, err := json.Marshal(&subject)
			if err != nil {
				log.Print("Bad json... SAD!")
				break
			}

			message := []byte(json)

			err = conn.WriteMessage(websocket.TextMessage, message)
			messagesSincePrint++

			subject.GenerateUpdate(currentTime)

			if err != nil {
				log.Print("Write ERR: ", err)
				break
			}
		}
	}
}

func onAllInstrumentsUpdate(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Print("Upgrade ERR: ", err)
	}

	defer conn.Close()

	log.Print(req.RemoteAddr, " subscribed to all instruments")

	for {
		currentTime := time.Now().UnixMilli()

		for _, ticker := range tickers {
			if ticker.NextUpdate < currentTime {
				json, err := json.Marshal(&ticker)
				if err != nil {
					log.Print("Bad json... SAD!")
					break
				}

				message := []byte(json)

				err = conn.WriteMessage(websocket.TextMessage, message)
				messagesSincePrint++

				ticker.GenerateUpdate(currentTime)

				if err != nil {
					log.Print("Write ERR: ", err)
					return
				}
			}
		}
	}
}

func monitor() {
	for range time.Tick(time.Second * 10) {
		log.Print(messagesSincePrint, " messages last 15 sec (", messagesSincePrint/15, " msg/s)")
		messagesSincePrint = 0
	}
}

func main() {
	rand.Seed(SEED)

	// Need to create the tickers AFTER seeding rand
	tickers = ticker.CreateAllTickers()

	http.HandleFunc("/stocks", onAllInstrumentsUpdate)
	http.HandleFunc("/stock/", onInstrumentUpdate)

	go monitor()

	if err := http.ListenAndServe(":4503", nil); err != nil {
		log.Fatal(err)
	}
}
