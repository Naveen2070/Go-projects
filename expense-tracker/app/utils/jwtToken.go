package utilities

import (
	"ExpenseTracker/app/model"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(data model.User) (string, error) {
	err := LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set the token expiration time (e.g., 1 hour)
	expirationTime := time.Now().Add(24 * time.Hour).Unix()

	// Create the Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   data.ID,
		"username": data.Username,
		"exp":      expirationTime,
	})

	// Convert the JWT_SECRET to []byte
	secret := []byte(os.Getenv("JWT_SECRET"))

	t, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyToken(tokenString string) (bool, error) {
	err := LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
