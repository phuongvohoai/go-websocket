package main

import (
	"log"
	"net/http"

	"phuong/go-websocket/routes"
	"phuong/go-websocket/ws"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ws.NewChatHub()
	go ws.ChatHub.Run()

	routes.MapWebRoutes()
	routes.MapWsRoutes()

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
