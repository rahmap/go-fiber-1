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

	apiV1 := app.Group("/api/v1")

	user.RoutesV1(apiV1)

	err := app.Listen(":5000")
	if err != nil {
		return
	}
}
