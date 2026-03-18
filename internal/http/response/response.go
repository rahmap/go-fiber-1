package response

import "github.com/gofiber/fiber/v2"

type Envelope struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorEnvelope struct {
	Status  bool   `json:"status" example:"false"`
	Message string `json:"message" example:"invalid user id"`
}

func Success(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(Envelope{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(ErrorEnvelope{
		Status:  false,
		Message: message,
	})
}
