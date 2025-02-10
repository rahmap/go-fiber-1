package user

import (
	"fiber-rest-api/database"
	"fiber-rest-api/module/user/model"
	"fiber-rest-api/module/user/structs"
	"github.com/gofiber/fiber/v2"
)

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

func GetUsers(c *fiber.Ctx) error {
	var users []model.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": result.Error.Error(),
		})
	}

	return c.JSON(structs.UsersResponse{
		Status:  true,
		Message: "success",
		Data:    users,
	})
}
