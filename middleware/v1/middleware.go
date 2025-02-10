package v1

import "github.com/gofiber/fiber/v2"

func ValidateHeader(c *fiber.Ctx) error {
	token := c.Get("X-TOKEN")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Missing X-TOKEN header",
		})
	}

	if token != "your-secret-token" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid X-TOKEN",
		})
	}

	return c.Next()
}
