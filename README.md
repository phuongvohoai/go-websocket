# Go WebSocket Chat 💬 🚀

A real-time WebSocket chat application built with Go and the Gorilla WebSocket library. This project demonstrates how to implement a WebSocket server with client management, authentication, and real-time messaging.

## Features ✨

- **Real-time Messaging** 📨 - Instant message delivery between connected clients
- **User Authentication** 🔐 - API key-based connection validation
- **User Notifications** 📢 - System messages for user join/leave events
- **Clean UI** 🎨 - Modern interface built with TailwindCSS
- **Connection Status** 📡 - Visual indicators for connection state

### In Progress 🚧

- **Binary Data Transmission** 📁 - Support for sending binary data like files
- **Compression** 🗜️ - Message compression for optimized data transfer
- **Error Handling and Reconnection** 🔄 - Advanced error handling with automatic reconnection

## Installation 📥

```bash
# Install dependencies
go mod download

# Run the server
go run main.go
```

## Usage 🖥️

1. Create a `.env` file with your API key:
```
API_KEY=your_api_key_here
```

2. Start the server:
```bash
go run main.go
```

3. Open your browser and navigate to:
```
http://localhost:8080
```

4. Chat with other connected clients in real-time!

## Project Structure 🏗️

```
├── main.go           # Server entry point
├── routes/           # HTTP and WebSocket routes
│   ├── web.go        # Web routes
│   └── ws.go         # WebSocket routes
├── ws/               # WebSocket implementation
│   ├── client.go     # Client connection handling
│   ├── hub.go        # Central hub for managing clients
│   └── message.go    # Message structure and utilities
└── web/              # Frontend assets
    └── index.html    # Chat interface
```