package main

import (
	"log"
	"micro-ocr-service/api/routes"
	"os"

	"micro-ocr-service/api/presenter"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/otiai10/gosseract/v2"
)

func main() {

	// load .env variables
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file or .env has missing")
	}

	// connect database mysql
	// sqlEnable, _ := strconv.ParseBool(os.Getenv("DB_SQL_ENABLE"))

	// if sqlEnable {
	// 	settings.ConnectMySqlDb()
	// }

	tesseract := gosseract.NewClient()

	app := fiber.New()
	app.Use(cors.New())

	// list of routes
	routes.SetupIndexRouter(app.Group("/"))
	routes.SetupBookRouter(app.Group("/api/books"))

	routes.SetupOcrRouter(app.Group("/api/ocr/ktp"), tesseract)

	// custom routes
	app.Use(func(c *fiber.Ctx) error {

		// custom not found 404
		return c.Status(fiber.StatusNotFound).JSON(presenter.NotFoundResponse())
	})

	appPort := os.Getenv("APP_PORT")
	log.Fatal(app.Listen(":" + appPort))

}
