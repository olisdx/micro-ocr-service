package routes

import (
	"micro-ocr-service/api/presenter"

	"github.com/gofiber/fiber/v2"
)

func SetupIndexRouter(router fiber.Router) {

	router.Get("/", func(c *fiber.Ctx) error {
		// return c.Send([]byte("Welcome to Singapay End-User Service API!"))
		return c.JSON(presenter.WelcomeResponse())
	})
}
