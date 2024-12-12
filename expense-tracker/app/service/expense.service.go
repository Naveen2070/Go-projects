package service

import (
	connection "ExpenseTracker/app/db"
	"ExpenseTracker/app/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var expensesConnection = connection.ConnectDB()

type ExpenseService struct {
	db *gorm.DB
}

func NewExpenseService() *ExpenseService {
	return &ExpenseService{
		db: expensesConnection,
	}
}

func (s *ExpenseService) GetAllExpenses(userID uuid.UUID) ([]model.Expense, error) {
	var user model.User
	result := s.db.Preload("Expenses").First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return user.Expenses, nil
}

func (s *ExpenseService) GetExpenseByID(id uuid.UUID) (model.Expense, error) {
	var expense model.Expense
	result := s.db.First(&expense, id)
	if result.Error != nil {
		return model.Expense{}, result.Error
	}
	return expense, nil
}

func (s *ExpenseService) CreateExpense(expenses model.ExpensePayload) (bool, error) {
	parsedTime, _ := time.Parse("2006-01-02", expenses.Date)
	result := s.db.Create(&model.Expense{
		ID:          uuid.New(),
		Description: expenses.Description,
		UserID:      expenses.UserID,
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
	result := s.db.First(&expense, id)
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

	result = s.db.Model(&expense).Updates(&expenseToUpdate)
	if result.Error != nil {
		return model.Expense{}, result.Error
	}
	return expense, nil
}

func (s *ExpenseService) DeleteExpense(id uuid.UUID) error {
	var expense model.Expense
	result := s.db.First(&expense, id)
	if result.Error != nil {
		return result.Error
	}

	result = s.db.Delete(&expense)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
