package service

import (
	connection "ExpenseTracker/app/db"
	"ExpenseTracker/app/model"
	"time"

	"github.com/google/uuid"
)

var expensesConnection = connection.ConnectDB()

type ExpenseService struct{}

func NewExpenseService() *ExpenseService {
	return &ExpenseService{}
}

func (s *ExpenseService) GetAllExpenses() ([]model.Expense, error) {
	var expenses []model.Expense
	result := expensesConnection.Find(&expenses)
	if result.Error != nil {
		return nil, result.Error
	}
	return expenses, nil
}

func (s *ExpenseService) GetExpenseByID(id uuid.UUID) (model.Expense, error) {
	var expense model.Expense
	result := expensesConnection.First(&expense, id)
	if result.Error != nil {
		return model.Expense{}, result.Error
	}
	return expense, nil
}

func (s *ExpenseService) CreateExpense(expenses model.ExpensePayload) (bool, error) {
	parsedTime, _ := time.Parse("2006-01-02", expenses.Date)
	result := expensesConnection.Create(&model.Expense{
		ID:          uuid.New(),
		Description: expenses.Description,
		Amount:      expenses.Amount,
		Category:    expenses.Category,
		Date:        parsedTime,
	})
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (s *ExpenseService) UpdateExpense(id uuid.UUID, updatedExpense model.ExpensePayload) (model.Expense, error) {
	var expense model.Expense
	result := expensesConnection.First(&expense, id)
	if result.Error != nil {
		return model.Expense{}, result.Error
	}
	parsedTime, _ := time.Parse("2006-01-02", updatedExpense.Date)
	expenseToUpdate := model.Expense{
		Description: updatedExpense.Description,
		Amount:      updatedExpense.Amount,
		Category:    updatedExpense.Category,
		Date:        parsedTime,
		UpdatedAt:   time.Now(),
	}

	result = expensesConnection.Model(&expense).Updates(&expenseToUpdate)
	if result.Error != nil {
		return model.Expense{}, result.Error
	}
	return expense, nil
}

func (s *ExpenseService) DeleteExpense(id uuid.UUID) error {
	var expense model.Expense
	result := expensesConnection.First(&expense, id)
	if result.Error != nil {
		return result.Error
	}

	result = expensesConnection.Delete(&expense)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
