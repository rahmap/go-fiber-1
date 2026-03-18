package user

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router, handler *Handler) {
	users := router.Group("/users")
	users.Get("/", handler.GetAll)
	users.Get("/:id", handler.GetByID)
	users.Post("/", handler.Create)
}
