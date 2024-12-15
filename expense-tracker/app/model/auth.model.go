package model

import "github.com/google/uuid"

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TwoFactorPayload struct {
	UserId uuid.UUID `json:"userId"`
	Code   string    `json:"code"`
}
