package routes

import (
	"micro-ocr-service/api/handlers"
	"micro-ocr-service/pkg/service"

	"github.com/gofiber/fiber/v2"
	"github.com/otiai10/gosseract/v2"
)

func SetupOcrRouter(router fiber.Router, client *gosseract.Client) {

	service := service.NewOcrService()
	handler := handlers.NewOcrHandler(service, client)

	router.Post("/", handler.SendKTP)

}
