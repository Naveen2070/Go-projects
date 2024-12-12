package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primarykey"`
	Username  string
	Email     string
	Password  string
	Expenses  []Expense `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "users"
}
