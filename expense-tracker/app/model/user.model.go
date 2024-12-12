package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Expenses  []Expense  `json:"expenses"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type UserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
