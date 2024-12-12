package migrations

import (
	"ExpenseTracker/app/db/schema"
	"errors"

	"gorm.io/gorm"
)

func ExcuteMigrations(db *gorm.DB) error {
	var err error
	// Ensure users table is created before expenses table
	err = userMigrations(db)

	// Ensure expenses table is created
	err = expenseMigrations(db)

	return err

}

func userMigrations(db *gorm.DB) error {

	err := db.AutoMigrate(&schema.User{})
	if err != nil {
		return errors.New("failed to migrate users table: " + err.Error())
	}

	return nil
}

func expenseMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&schema.Expense{})
	if err != nil {
		return errors.New("failed to migrate expenses table: " + err.Error())
	}

	return nil
}
