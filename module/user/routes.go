package user

import "github.com/gofiber/fiber/v2"

func RoutesV1(router fiber.Router) {
	userGroup := router.Group("/user")
	userGroup.Get("/find", GetUser)
	userGroup.Get("/all", GetUsers)
}
