package main

import (
	"ExpenseTracker/app"
	utilities "ExpenseTracker/app/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

var PORT string

// init initializes the PORT variable by retrieving the value from the
// environment variable "PORT". If the environment variable is not set,
// it defaults to "8080".
func init() {
	err := utilities.LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT = os.Getenv("SERVER_PORT")
	if PORT == "" {
		PORT = "8080"
	}
}

// main is the entry point for the program.
//
// It sets up a Fiber HTTP server and configures a single route
// at "/". The route responds with "Hello, World!".
//
// The server is then started and left to run until it is manually
// stopped.
func main() {
	appInstance := fiber.New()

	app.Setup(appInstance)

	log.Println("Server is running on port", PORT)
	log.Fatal(appInstance.Listen(":" + PORT))
}
