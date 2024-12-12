package utilities

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type PassFactory struct{}

func NewPassFactory() *PassFactory {
	return &PassFactory{}
}

func (p *PassFactory) GeneratePassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

func (p *PassFactory) ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
