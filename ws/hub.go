package ws

import (
	"fmt"
	"log"
)

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

var ChatHub *Hub

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte, 10),
	}
}

func NewChatHub() {
	ChatHub = NewHub()
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			message := NewMessageEvent("system", fmt.Sprintf("Welcome %s 🫡", client.name))
			h.broadcast <- message.ToJson()
			// go func() {
			// 	h.broadcast <- []byte("Client joined!")
			// }()
			log.Printf("Client registered: %v", client.connection.RemoteAddr())

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				message := NewMessageEvent("system", fmt.Sprintf("Goodbye %s 👋", client.name))
				h.broadcast <- message.ToJson()
				delete(h.clients, client)
				close(client.send)
				log.Printf("Client unregistered: %v", client.connection.RemoteAddr())
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
					log.Printf("Sent message to client: %v", client.connection.RemoteAddr())
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) RegisterClient(client *Client) {
	h.register <- client
}
