package app

import (
	ExpenseController "ExpenseTracker/app/controller"
	"ExpenseTracker/app/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Use(logger.Logger())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	ExpenseController.RegisterRoutes(app)

}
