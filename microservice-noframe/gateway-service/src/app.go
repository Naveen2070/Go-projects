package app

import (
	"encoding/json"
	"fmt"
	"gateway/src/utils"
	"net/http"
)

func App(server *http.Server) {
	handler := http.NewServeMux()

	handler.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		var request map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := utils.ForwardRequest("localhost:9001", request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	handler.HandleFunc("/api/posts", func(w http.ResponseWriter, r *http.Request) {
		var request map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := utils.ForwardRequest("localhost:9002", request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	fmt.Println("Gateway Service is running on :8000")

	// Assign the handler to the server
	server.Handler = handler

	// Start the server with the configured handler
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Could not listen on %s: %v\n", server.Addr, err)
	}
}
