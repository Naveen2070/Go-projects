package utils

import (
	"encoding/json"
	"io"
	"net"
)

func ForwardRequest(address string, request map[string]interface{}) ([]byte, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	requestData, _ := json.Marshal(request)
	conn.Write(requestData)

	response, err := io.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	return response, nil
}
