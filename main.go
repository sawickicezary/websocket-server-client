package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

var counter = 0

type Server struct {
	clients map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (*Server) handleWS(ws *websocket.Conn) {
	// fmt.Println("New client requested")
	ws.Write([]byte(fmt.Sprintf("Hello there client number: %v", counter)))
	counter += 1
}

func main() {
	server := NewServer()
	http.Handle("/hello", websocket.Handler(server.handleWS))
	http.ListenAndServe(":23718", nil)
}
