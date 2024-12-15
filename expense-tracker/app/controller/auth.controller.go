package controller

import (
	"ExpenseTracker/app/middleware"
	"ExpenseTracker/app/model"
	"ExpenseTracker/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var authService *service.AuthService

func RegisterAuthRoutes(api fiber.Router) {
	// Initialize the service
	authService = service.NewAuthService(service.NewUserService())

	api.Post("/register", registerUser)
	api.Post("/login", loginUser)
	api.Get("/initialize2fa/:id", middleware.JWTProtected, initialize2fa)
	api.Post("/verify2fa", middleware.JWTProtected, verify2fa)
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

func initialize2fa(c *fiber.Ctx) error {
	unParsedID := c.Params("id")

	id, err := uuid.Parse(unParsedID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}

	result, err := authService.InitializeTwoFactorAuth(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to initialize 2FA", "error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "2FA initialized successfully", "qrCode": result.QRCode, "url": result.URL})
}

func verify2fa(c *fiber.Ctx) error {
	var user model.TwoFactorPayload
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request body"})
	}

	result, err := authService.VerifyTwoFactorAuth(user.UserId, user.Code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to verify 2FA", "error": err.Error()})
	}

	if !result {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Invalid code provided"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "2FA verified successfully"})
}
