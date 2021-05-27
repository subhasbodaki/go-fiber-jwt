package main

import (
	"fmt"
	"log"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is up and running")
	})

	app.Post("/login", login)

	app.Get("/hello", authRequired(), func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token) // get the token details (decode token)
		claims := user.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		return c.SendString(fmt.Sprintf("Hello user with email %s", email))
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

func authRequired() func(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SigningKey: []byte("secret"),
	})
}

func login(c *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user request
	err := c.BodyParser(&user) //It recieves the login details from Body and parse it
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse json",
		})
	}

	//match email & password
	if user.Email != "sbodaki@gmail.com" || user.Password != "password" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	//create token
	token := jwt.New(jwt.SigningMethodHS256)

	//set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = "sbodaki@gmail.com"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7) //Week

	//generate encoded token
	s, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	//send encoded token as response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
	})

}
