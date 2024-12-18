package main

import (
	app "gateway/src"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":8000",
		Handler: nil, // Set to nil, will be set in app.App
	}

	// Initialize the app with the server
	app.App(server)
}
