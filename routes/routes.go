package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subhasbodaki/go-fiber-jwt/handler"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.Server)

	app.Post("/login", handler.Login)

	app.Get("/me", handler.Hello)
}
