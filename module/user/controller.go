package user

import "github.com/gofiber/fiber/v2"
import "fiber-rest-api/module/user/structs"

func GetUser(c *fiber.Ctx) error {
	var request structs.FindByIdRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid request body",
		})
	}

	if request.ID == nil || *request.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "ID is required and must be greater than 0",
		})
	}

	return c.JSON(structs.UserResponse{
		Status:  true,
		Message: "success",
		Data: structs.UserDataResponse{
			ID:      request.ID,
			Name:    "Rahma AP",
			Address: "Yogyakarta",
		},
	})
}
