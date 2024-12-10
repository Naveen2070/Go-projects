package ExpenseController

import (
	models "ExpenseTracker/app/model"
	ExpenseService "ExpenseTracker/app/service"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	api.Get("/expenses", getAllExpenses)
	api.Get("/expenses/:id", getExpenseByID)
	api.Post("/expenses", createExpense)
	api.Put("/expenses/:id", updateExpense)
	api.Delete("/expenses/:id", deleteExpense)
}

func getAllExpenses(c *fiber.Ctx) error {
	expenses, err := ExpenseService.GetAllExpenses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to retrieve expenses"})
	}
	return c.JSON(expenses)
}

func getExpenseByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid expense ID"})
	}
	expense, err := ExpenseService.GetExpenseByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "expense not found"})
	}
	return c.JSON(expense)
}

func createExpense(c *fiber.Ctx) error {
	var expense models.CreateExpenseRequest
	if err := c.BodyParser(&expense); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	newExpense := ExpenseService.CreateExpense(expense)
	return c.JSON(newExpense)
}

func updateExpense(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid expense ID"})
	}
	var updatedExpense models.Expense
	if err := c.BodyParser(&updatedExpense); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	updatedExpense, err = ExpenseService.UpdateExpense(id, updatedExpense)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "expense not found"})
	}
	return c.JSON(updatedExpense)
}

func deleteExpense(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid expense ID"})
	}
	if err := ExpenseService.DeleteExpense(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "expense not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
