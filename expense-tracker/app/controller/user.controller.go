package controller

import (
	"ExpenseTracker/app/model"
	"ExpenseTracker/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var userService *service.UserService

func RegisterUserRoutes(api fiber.Router) {
	// Initialize the service
	userService = service.NewUserService()

	api.Get("/", getAllUsers)
	api.Get("/:id", getUserByID)
	api.Post("/", createUser)
	api.Put("/:id", updateUser)
	api.Delete("/:id", deleteUser)
	api.Put("/updatePassword/:id", updatePassword)
}

func getAllUsers(c *fiber.Ctx) error {
	users, err := userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to retrieve users"})
	}
	return c.JSON(fiber.Map{"data": users})
}

func getUserByID(c *fiber.Ctx) error {
	unParsedID := c.Params("id")

	id, err := uuid.Parse(unParsedID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}

	user, err := userService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user with id " + unParsedID + " not found", "data": []string{}})
	}
	return c.JSON(fiber.Map{"data": &[]model.User{user}})
}

func createUser(c *fiber.Ctx) error {
	var user model.UserPayload
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request body"})
	}
	result, err := userService.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to create user", "error": err.Error()})
	}
	if !result {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to create user"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func updateUser(c *fiber.Ctx) error {
	unParsedID := c.Params("id")

	id, err := uuid.Parse(unParsedID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}

	var updatedUser model.UserPayload
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request body"})
	}

	user, err := userService.UpdateUser(id, updatedUser)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
	}
	return c.JSON(fiber.Map{"data": &[]model.User{user}})
}

func deleteUser(c *fiber.Ctx) error {
	unParsedID := c.Params("id")

	id, err := uuid.Parse(unParsedID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}

	if err := userService.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func updatePassword(c *fiber.Ctx) error {
	unParsedID := c.Params("id")

	id, err := uuid.Parse(unParsedID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}

	if err := userService.UpdatePassword(id, c.FormValue("password")); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
