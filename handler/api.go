package handler

import "github.com/gofiber/fiber/v2"

func Server(c *fiber.Ctx) error {
	return c.SendString("Server is up and running")
}
