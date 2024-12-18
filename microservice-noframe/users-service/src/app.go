package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type Post struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

var posts = []Post{{ID: 1, UserID: 1, Content: "Hello World!"}}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

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
	var response []byte

	switch action {
	case "create":
		postData := request["data"].(map[string]interface{})
		newPost := Post{
			ID:      int(postData["id"].(float64)),
			UserID:  int(postData["user_id"].(float64)),
			Content: postData["content"].(string),
		}
		posts = append(posts, newPost)
		response = []byte("Post created successfully")
	case "read":
		response, _ = json.Marshal(posts)
	case "update":
		postData := request["data"].(map[string]interface{})
		id := int(postData["id"].(float64))
		for i, post := range posts {
			if post.ID == id {
				posts[i].Content = postData["content"].(string)
				response = []byte("Post updated successfully")
				break
			}
		}
	case "delete":
		id := int(request["data"].(map[string]interface{})["id"].(float64))
		for i, post := range posts {
			if post.ID == id {
				posts = append(posts[:i], posts[i+1:]...)
				response = []byte("Post deleted successfully")
				break
			}
		}
	default:
		response = []byte("Unknown action")
	}

	conn.Write(response)
}
