package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var loggerConfig = logger.Config{
	Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
}

func Logger() fiber.Handler {
	return logger.New(loggerConfig)
}
