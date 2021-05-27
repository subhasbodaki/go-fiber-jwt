package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/subhasbodaki/go-fiber-jwt/handler"
	"github.com/subhasbodaki/go-fiber-jwt/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use("/me", handler.AuthRequired())

	routes.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
