package main

import (
	"fmt"
	"net"
	app "posts/src"
)

func main() {
	address := "localhost:9002"

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error starting Posts Service:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Posts Service is listening on", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}

		go app.HandleConnection(conn)
	}
}
