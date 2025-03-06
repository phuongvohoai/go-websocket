package main

import "github.com/gorilla/websocket"

type Client struct {
	connection *websocket.Conn
	hub        *Hub
	send       chan []byte
}

func NewClient(connection *websocket.Conn, hub *Hub) *Client {
	return &Client{
		connection: connection,
		hub:        hub,
		send:       make(chan []byte),
	}
}

func (c *Client) ReadMessage() {
	defer func() {
		c.hub.unregister <- c
		c.connection.Close()
	}()

	for {
		_, msg, err := c.connection.ReadMessage()
		if err != nil {
			break
		}

		c.hub.broadcast <- msg
	}
}

func (c *Client) WriteMessage() {
	defer func() {
		c.connection.Close()
	}()

	for msg := range c.send {
		c.connection.WriteMessage(websocket.TextMessage, msg)
	}

	c.connection.WriteMessage(websocket.CloseMessage, []byte{})
}
