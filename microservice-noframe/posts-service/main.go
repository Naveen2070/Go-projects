package main

import (
	"fmt"
	"net"
	app "posts/src" // Make sure to replace this with the actual import path of your app package
)

func main() {
	address := "localhost:9001"

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error starting User Service:", err)
		return
	}
	defer listener.Close()

	fmt.Println("User Service is listening on", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}

		go app.HandleConnection(conn)
	}
}
