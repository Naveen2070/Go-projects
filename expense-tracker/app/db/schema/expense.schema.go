package expenseSchema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Expense struct {
	ID          uuid.UUID `gorm:"primarykey"`
	Description string
	Amount      float64
	Category    string
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (Expense) TableName() string {
	return "expenses"
}
