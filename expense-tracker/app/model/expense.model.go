package expenseModel

import "time"

type Expense struct {
	ID          uint       `json:"id"`
	Description string     `json:"description"`
	Amount      float64    `json:"amount"`
	Category    string     `json:"category"`
	Date        string     `json:"date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type CreateExpenseRequest struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Date        string  `json:"date"`
}
