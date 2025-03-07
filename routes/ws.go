package routes

import (
	"log"
	"net/http"
	"os"

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

	// Handle authentication
	apiKey := r.URL.Query()["api_key"]
	if len(apiKey) != 1 || apiKey[0] != os.Getenv("API_KEY") {
		invalidKeyMessage := ws.NewMessageEvent("system", "Invalid API key")
		connection.WriteMessage(websocket.TextMessage, invalidKeyMessage.ToJson())
		log.Println("Invalid API key")
		connection.Close()
		return
	}

	var msg ws.MessageEvent
	err = connection.ReadJSON(&msg)
	if err != nil || msg.Type != "register" || msg.User == "" {
		invalidRegisterMessage := ws.NewMessageEvent("system", "Invalid registration message")
		connection.WriteMessage(websocket.TextMessage, invalidRegisterMessage.ToJson())
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
