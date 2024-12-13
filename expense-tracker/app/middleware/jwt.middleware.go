package middleware

import (
	utilities "ExpenseTracker/app/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
)

func JWTProtected(c *fiber.Ctx) error {
	err := utilities.LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ContextKey: "jwt",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Return status 401 and failed authentication error.
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		},
	})(c)
}
