package app

import (
	"ExpenseTracker/app/controller"
	"ExpenseTracker/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Use(middleware.Logger())
	app.Use(middleware.Helmet())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//version group
	apiV1 := app.Group("/api/v1")

	//Path group
	authGroup := apiV1.Group("/auth")
	expenseGroup := apiV1.Group("/expenses")
	userGroup := apiV1.Group("/users")

	//add route specific middleware
	expenseGroup.Use(middleware.JWTProtected)
	userGroup.Use(middleware.JWTProtected)

	//controller registration
	controller.RegisterAuthRoutes(authGroup)
	controller.RegisterExpenseRoutes(expenseGroup)
	controller.RegisterUserRoutes(userGroup)

}
