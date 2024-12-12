package controller

import (
	"ExpenseTracker/app/model"
	"ExpenseTracker/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var expenseService *service.ExpenseService

func RegisterExpenseRoutes(api fiber.Router) {
	// Initialize the service
	expenseService = service.NewExpenseService()

	api.Get("/:userID", getAllExpenses)
	api.Get("/getExpenseById/:id", getExpenseByID)
	api.Post("/", createExpense)
	api.Put("/:id", updateExpense)
	api.Delete("/:id", deleteExpense)
}

func getAllExpenses(c *fiber.Ctx) error {
	unParsedID := c.Params("userID")

	id, err := uuid.Parse(unParsedID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}

	expenses, err := expenseService.GetAllExpenses(id)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No expenses found for user with id " + unParsedID, "data": []string{}})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to retrieve expenses", "error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": expenses})
}

func getExpenseByID(c *fiber.Ctx) error {
	unParsedID := c.Params("id")

	id, err := uuid.Parse(unParsedID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}

	expense, err := expenseService.GetExpenseByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "expense with id " + unParsedID + " not found", "data": []string{}})
	}
	return c.JSON(fiber.Map{"data": &[]model.Expense{expense}})
}

func createExpense(c *fiber.Ctx) error {
	var expenses model.ExpensePayload
	if err := c.BodyParser(&expenses); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request body"})
	}
	result, _ := expenseService.CreateExpense(expenses)
	if !result {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to create expense"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Expense created successfully"})
}

func updateExpense(c *fiber.Ctx) error {
	unParsedID := c.Params("id")

	id, err := uuid.Parse(unParsedID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}

	var updatedExpense model.ExpensePayload
	if err := c.BodyParser(&updatedExpense); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request body"})
	}

	expense, err := expenseService.UpdateExpense(id, updatedExpense)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "expense not found"})
	}
	return c.JSON(fiber.Map{"data": &[]model.Expense{expense}})
}

func deleteExpense(c *fiber.Ctx) error {
	unParsedID := c.Params("id")

	id, err := uuid.Parse(unParsedID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid UUID"})
	}

	if err := expenseService.DeleteExpense(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "expense not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
