package ExpenseService

import (
	model "ExpenseTracker/app/model"
	"errors"
)

func GetAllExpenses() []model.Expense {
	return model.Expenses
}

func GetExpenseByID(id int) (model.Expense, error) {
	for _, expense := range model.Expenses {
		if expense.ID == id {
			return expense, nil
		}
	}
	return model.Expense{}, errors.New("expense not found")
}

func CreateExpense(expense model.Expense) model.Expense {
	expense.ID = len(model.Expenses) + 1
	model.Expenses = append(model.Expenses, expense)
	return expense
}

func UpdateExpense(id int, updatedExpense model.Expense) (model.Expense, error) {
	for i, expense := range model.Expenses {
		if expense.ID == id {
			model.Expenses[i] = updatedExpense
			model.Expenses[i].ID = id
			return model.Expenses[i], nil
		}
	}
	return model.Expense{}, errors.New("expense not found")
}

func DeleteExpense(id int) error {
	for i, expense := range model.Expenses {
		if expense.ID == id {
			model.Expenses = append(model.Expenses[:i], model.Expenses[i+1:]...)
			return nil
		}
	}
	return errors.New("expense not found")
}
