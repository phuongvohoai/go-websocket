package main

import (
	"log"
	"net/http"

	"phuong/go-websocket/routes"
	"phuong/go-websocket/ws"
)

func main() {
	ws.NewChatHub()
	go ws.ChatHub.Run()

	routes.MapWebRoutes()
	routes.MapWsRoutes()

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
