package main

import (
	"fiber-rest-api/module/user"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	app.Get("/get-user", user.GetUser)

	err := app.Listen(":5000")
	if err != nil {
		return
	}
}
