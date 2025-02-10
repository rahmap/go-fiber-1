package user

import "github.com/gofiber/fiber/v2"
import "fiber-rest-api/module/user/structs"

func GetUser(c *fiber.Ctx) error {
	return c.JSON(structs.UserResponse{
		Status:  true,
		Message: "success",
		Data: structs.UserData{
			ID:      1,
			Name:    "Rahma AP",
			Address: "Yogyakarta",
		},
	})
}
