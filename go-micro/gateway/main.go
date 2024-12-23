package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

func main() {
	// Define the HTTP server
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		// Read the message from the HTTP request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		message := strings.TrimSpace(string(body))
		if message == "" {
			http.Error(w, "Message cannot be empty", http.StatusBadRequest)
			return
		}

		// Forward the message to the TCP server
		response, err := sendToTCPServer(message)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to communicate with TCP server: %v", err), http.StatusInternalServerError)
			return
		}

		// Write the TCP server's response back to the HTTP client
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(response))
	})

	// Start the HTTP server
	fmt.Println("HTTP server is listening on http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf("HTTP server failed: %v\n", err)
	}
}

// sendToTCPServer connects to the TCP server, sends a message, and returns the response.
func sendToTCPServer(message string) (string, error) {
	// Connect to the TCP server
	conn, err := net.Dial("tcp", "localhost:8080")
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
