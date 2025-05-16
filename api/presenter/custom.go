package presenter

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func NotFoundResponse() *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  "Data or route not found!",
	}
}

func WelcomeResponse() *fiber.Map {
	appName := os.Getenv("APP_NAME")

	return &fiber.Map{
		"status": true,
		"data":   "Welcome to " + appName,
		"error":  nil,
	}
}
