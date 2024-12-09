package main

import (
	"log"
	"os"

	"ExpenseTracker/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

var PORT string

// init initializes the PORT variable by retrieving the value from the
// environment variable "PORT". If the environment variable is not set,
// it defaults to "8080".

func init() {
	PORT = os.Getenv("PORT")
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
	app := fiber.New()

	app.Use(logger.Logger())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":8080"))
}
