package src

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Handler interface for processing messages.
type Handler interface {
	ProcessMessage(input string) string
}

// ListenAndServe starts the TCP server.
func ListenAndServe(port int, handler Handler) error {
	listener, err := net.Listen("tcp", ":"+fmt.Sprint(port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", port, err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("failed to accept connection: %v\n", err)
			continue
		}

		fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr().String())

		go startReadAndSend(conn, handler)
	}
}

// startReadAndSend continuously reads from and writes to the TCP connection.
func startReadAndSend(conn net.Conn, handler Handler) {
	defer conn.Close()
	fmt.Printf("New connection from %s\n", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	for {
		// Read the message from the client
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Connection from %s closed: %v\n", conn.RemoteAddr().String(), err)
			break
		}

		// Trim the message and process it
		message = strings.TrimSpace(message)
		fmt.Printf("Received: %s\n", message)

		// Get the response using the Handler
		response := handler.ProcessMessage(message)

		// Send the response back to the client
		_, err = conn.Write([]byte(response + "\n"))
		if err != nil {
			fmt.Printf("Failed to send response to %s: %v\n", conn.RemoteAddr().String(), err)
			break
		}
	}
}

// sendToTCPServer connects to the TCP server, sends a message, and returns the response.
func SendToTCPServer(Port int, message string) (string, error) {
	// Connect to the TCP server
	conn, err := net.Dial("tcp", "localhost:"+fmt.Sprint(Port))
	if err != nil {
		return "", fmt.Errorf("failed to connect to TCP server: %w", err)
	}
	defer conn.Close()

	// Send the message to the TCP server
	_, err = conn.Write([]byte(message + "\n"))
	if err != nil {
		return "", fmt.Errorf("failed to send message to TCP server: %w", err)
	}

	// Read the response from the TCP server
	responseBuffer := make([]byte, 1024)
	n, err := conn.Read(responseBuffer)
	if err != nil {
		return "", fmt.Errorf("failed to read response from TCP server: %w", err)
	}

	return string(responseBuffer[:n]), nil
}
