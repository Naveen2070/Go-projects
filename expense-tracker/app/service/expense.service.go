package ExpenseService

import (
	database "ExpenseTracker/app/db"
	model "ExpenseTracker/app/model"
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

func GetExpenseByID(id int) (model.Expense, error) {
	var expense model.Expense
	result := db.First(&expense, id)
	if result.Error != nil {
		return model.Expense{}, result.Error
	}
	return expense, nil
}

func CreateExpense(expense model.CreateExpenseRequest) model.CreateExpenseRequest {
	result := db.Create(&expense)
	if result.Error != nil {
		return model.CreateExpenseRequest{}
	}
	return expense
}

func UpdateExpense(id int, updatedExpense model.Expense) (model.Expense, error) {
	var expense model.Expense
	result := db.First(&expense, id)
	if result.Error != nil {
		return model.Expense{}, result.Error
	}

	updatedExpense = model.Expense{
		Description: updatedExpense.Description,
		Amount:      updatedExpense.Amount,
		Category:    updatedExpense.Category,
		Date:        updatedExpense.Date,
		UpdatedAt:   updatedExpense.UpdatedAt,
	}

	result = db.Save(&updatedExpense)
	if result.Error != nil {
		return model.Expense{}, result.Error
	}
	return expense, nil
}

func DeleteExpense(id int) error {
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
