package routes

import (
	"log"
	"net/http"

	"phuong/go-websocket/ws"

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

	var msg ws.MessageEvent
	err = connection.ReadJSON(&msg)
	if err != nil || msg.Type != "register" || msg.User == "" {
		log.Println("Invalid registration message", msg)
		connection.Close()
		return
	}

	client := ws.NewClient(connection, ws.ChatHub, msg.User)
	ws.ChatHub.RegisterClient(client)
	go client.ReadMessage()
	go client.WriteMessage()
}

func MapWsRoutes() {
	http.HandleFunc("/ws", serveWebSocket)
}
