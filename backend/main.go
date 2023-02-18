package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func createWSHandleFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	for {
		currentTime := time.Now().UnixMicro()
		message := []byte(strconv.Itoa(int(currentTime)))

		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	http.HandleFunc("/updates/first", createWSHandleFunc)
	http.HandleFunc("/updates/second", createWSHandleFunc)

	if err := http.ListenAndServe(":4503", nil); err != nil {
		log.Fatal(err)
	}
}
