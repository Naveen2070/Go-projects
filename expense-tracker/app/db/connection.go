package connection

import (
	"ExpenseTracker/app/db/migrations"
	utilities "ExpenseTracker/app/utils"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	HOST string
	USER string
	PASS string
	DB   string
	PORT string
)

// init initializes the database connection variables by retrieving the values from the
// environment variables "DB_HOST", "DB_USER", "DB_PASS", "DB_NAME", and "DB_PORT". If the
// environment variables are not set, it defaults to "localhost", "postgres", "postgres",
// "go_dev", and "5432", respectively.
func init() {
	err := utilities.LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	HOST = os.Getenv("DB_HOST")
	USER = os.Getenv("DB_USER")
	PASS = os.Getenv("DB_PASS")
	DB = os.Getenv("DB_NAME")
	PORT = os.Getenv("DB_PORT")

	if HOST == "" {
		HOST = "localhost"
	}
	if USER == "" {
		USER = "postgres"
	}
	if PASS == "" {
		PASS = "postgres"
	}
	if DB == "" {
		DB = "go_dev"
	}
	if PORT == "" {
		PORT = "5432"
	}
}

func ConnectDB() *gorm.DB {
	dsn := "host=" + HOST + " user=" + USER + " password=" + PASS + " dbname=" + DB + " port= " + PORT
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	//run migrations
	err = migrations.ExcuteMigrations(db)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
