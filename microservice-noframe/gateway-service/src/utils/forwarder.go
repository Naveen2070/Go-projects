package utils

import (
	"encoding/json"
	"io"
	"net"
	"time"
)

func ForwardRequest(address string, request map[string]interface{}) ([]byte, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Set a deadline for the connection
	conn.SetDeadline(time.Now().Add(5 * time.Second))

	requestData, _ := json.Marshal(request)
	_, err = conn.Write(requestData)
	if err != nil {
		return nil, err
	}

	response, err := io.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	return response, nil
}
