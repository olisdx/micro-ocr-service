package routes

import (
	"micro-ocr-service/api/handlers"
	"micro-ocr-service/pkg/service"

	"github.com/gofiber/fiber/v2"
)

func SetupBookRouter(router fiber.Router) {

	service := service.NewService()

	router.Get("/", handlers.GetBooks(service))
	router.Post("/", handlers.AddBook(service))
	router.Patch("/", handlers.UpdateBook(service))
	router.Delete("/", handlers.RemoveBook(service))

}
