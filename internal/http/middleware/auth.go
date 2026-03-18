package middleware

import "github.com/gofiber/fiber/v2"

func ValidateToken(expectedToken string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("X-TOKEN")

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  false,
				"message": "missing X-TOKEN header",
			})
		}

		if token != expectedToken {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status":  false,
				"message": "invalid X-TOKEN",
			})
		}

		return c.Next()
	}
}
