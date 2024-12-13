package controller

import (
	"ExpenseTracker/app/model"
	"ExpenseTracker/app/service"

	"github.com/gofiber/fiber/v2"
)

var authService *service.AuthService

func RegisterAuthRoutes(api fiber.Router) {
	// Initialize the service
	authService = service.NewAuthService(service.NewUserService())

	api.Post("/register", registerUser)
	api.Post("/login", loginUser)
}

func registerUser(c *fiber.Ctx) error {
	var user model.UserPayload
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request body"})
	}

	result, err := authService.Register(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to register user"})
	}

	if !result {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to register user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func loginUser(c *fiber.Ctx) error {
	var user model.AuthPayload
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request body"})
	}

	result, err := authService.Login(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to login user", "error": err.Error()})
	}

	if result == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to generate token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User logged in successfully", "token": result})
}
