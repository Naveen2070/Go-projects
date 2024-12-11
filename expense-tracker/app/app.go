package app

import (
	"ExpenseTracker/app/controller"
	"ExpenseTracker/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Use(middleware.Logger())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//version group
	apiV1 := app.Group("/api/v1")

	//Path group
	expenseGroup := apiV1.Group("/expenses")

	//controller registration
	controller.RegisterExpenseRoutes(expenseGroup)

}
