package main

import (
	"fmt"
	"tcpconnect/src"
	"tcpconnect/src/handler"
)

func main() {
	// Create a new instance of the handler
	h := &handler.MessageHandler{}

	// Start the server
	err := src.ListenAndServe(8080, h)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
