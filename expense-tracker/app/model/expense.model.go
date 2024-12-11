package model

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID          uuid.UUID  `json:"id"`
	Description string     `json:"description"`
	Amount      float64    `json:"amount"`
	Category    string     `json:"category"`
	Date        time.Time  `json:"date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type ExpensePayload struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Date        string    `json:"date"`
}
