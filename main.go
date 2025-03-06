package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWebSocket(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	client := NewClient(connection, ChatHub)
	ChatHub.register <- client
	go client.ReadMessage()
	go client.WriteMessage()
}

func serveHomepage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	ChatHub = NewHub()
	go ChatHub.Run()

	http.HandleFunc("/ws", serveWebSocket)
	http.HandleFunc("/", serveHomepage)
	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
