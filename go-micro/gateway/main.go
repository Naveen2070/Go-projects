package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

type TCPMessage struct {
	Action  string      `json:"action"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	// Define the HTTP server
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Create the JSON structure
		message := TCPMessage{
			Action:  r.Method,
			Data:    json.RawMessage(body), // Forward raw data as JSON
			Message: "Request forwarded to TCP server",
		}

		// Convert the message to JSON
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			http.Error(w, "Failed to encode message as JSON", http.StatusInternalServerError)
			return
		}

		// Forward the JSON message to the TCP server
		response, err := SendToTCPServer(8080, string(jsonMessage))
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to communicate with TCP server: %v", err), http.StatusInternalServerError)
			return
		}

		// Write the TCP server's response back to the HTTP client
		w.Header().Set("Content-Type", "application/json")
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
