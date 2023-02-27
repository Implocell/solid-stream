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

var tickers []ticker.Ticker

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func onUpdatesRequest(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Print("Upgrade ERR: ", err)
	}

	defer conn.Close()

	symbol := req.URL.Path[9:len(req.URL.Path)]

	var subject *ticker.Ticker
	for _, ticker := range tickers {
		if ticker.Symbol == symbol {
			subject = &ticker
			break
		}
	}

	if (*subject == ticker.Ticker{}) {
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

			subject.GenerateUpdate(currentTime)

			if err != nil {
				log.Print("Write ERR: ", err)
				break
			}
		}
	}
}

func main() {
	rand.Seed(SEED)

	// Need to create the tickers AFTER seeding rand
	tickers = ticker.CreateAllTickers()

	http.HandleFunc("/updates/", onUpdatesRequest)

	if err := http.ListenAndServe(":4503", nil); err != nil {
		log.Fatal(err)
	}
}
