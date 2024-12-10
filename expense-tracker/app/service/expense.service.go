package ExpenseService

import (
	database "ExpenseTracker/app/db"
	model "ExpenseTracker/app/model"
	"time"

	"github.com/google/uuid"
)

var db = database.ConnectDB()

func GetAllExpenses() ([]model.Expense, error) {
	var expenses []model.Expense
	result := db.Find(&expenses)
	if result.Error != nil {
		return nil, result.Error
	}
	return expenses, nil
}

func GetExpenseByID(id uuid.UUID) (model.Expense, error) {
	var expense model.Expense
	result := db.First(&expense, id)
	if result.Error != nil {
		return model.Expense{}, result.Error
	}
	return expense, nil
}

func CreateExpense(expenses model.CreateExpenseRequest) (bool, error) {
	parsedTime, _ := time.Parse("2006-01-02", expenses.Date)
	result := db.Create(&model.Expense{
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

func UpdateExpense(id uuid.UUID, updatedExpense model.UpdateExpenseRequest) (model.Expense, error) {
	var expense model.Expense
	result := db.First(&expense, id)
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

	result = db.Model(&expense).Updates(&expenseToUpdate)
	if result.Error != nil {
		return model.Expense{}, result.Error
	}
	return expense, nil
}

func DeleteExpense(id uuid.UUID) error {
	var expense model.Expense
	result := db.First(&expense, id)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&expense)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
