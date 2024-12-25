package main

import (
	"fmt"

	"github.com/Naveen2070/Go-projects/go-micro/tcpconnect/src"
	"github.com/Naveen2070/Go-projects/go-micro/tcpconnect/src/handler"
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
