package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{{ID: 1, Name: "John Doe"}}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	// Read the incoming request
	requestData, err := io.ReadAll(conn)
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	var request map[string]interface{}
	if err := json.Unmarshal(requestData, &request); err != nil {
		fmt.Println("Invalid request format:", err)
		return
	}

	action := request["action"].(string)
	println(request)
	var response []byte

	switch action {
	case "create":
		userData := request["data"].(map[string]interface{})
		newUser := User{ID: int(userData["id"].(float64)), Name: userData["name"].(string)}
		users = append(users, newUser)
		response = []byte(`{"message": "User created successfully"}`)
	case "read":
		response, _ = json.Marshal(users)
	case "update":
		userData := request["data"].(map[string]interface{})
		id := int(userData["id"].(float64))
		for i, user := range users {
			if user.ID == id {
				users[i].Name = userData["name"].(string)
				response = []byte(`{"message": "User updated successfully"}`)
				break
			}
		}
	case "delete":
		id := int(request["data"].(map[string]interface{})["id"].(float64))
		for i, user := range users {
			if user.ID == id {
				users = append(users[:i], users[i+1:]...)
				response = []byte(`{"message": "User deleted successfully"}`)
				break
			}
		}
	default:
		response = []byte(`{"error": "Unknown action"}`)
	}

	// Send the response and close the connection
	conn.Write(response)
}
