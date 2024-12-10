package expenseSchema

import (
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	description string
	Amount      float64
	Category    string
	Date        string
}

func (Expense) TableName() string {
	return "expenses"
}
