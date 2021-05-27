package handler

import (
	"fmt"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token) // get the token details (decode token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	return c.SendString(fmt.Sprintf("Hello user with email %s", email))
}
