package user

import (
	"errors"
	"strconv"

	"fiber-rest-api/internal/http/response"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	users, err := h.service.GetAll()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "success", users)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid user id")
	}

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidUserID):
			return response.Error(c, fiber.StatusBadRequest, err.Error())
		case errors.Is(err, ErrUserNotFound):
			return response.Error(c, fiber.StatusNotFound, err.Error())
		default:
			return response.Error(c, fiber.StatusInternalServerError, err.Error())
		}
	}

	return response.Success(c, fiber.StatusOK, "success", user)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var request CreateRequest
	if err := c.BodyParser(&request); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}

	user, err := h.service.Create(request)
	if err != nil {
		switch {
		case errors.Is(err, ErrNameRequired):
			return response.Error(c, fiber.StatusBadRequest, err.Error())
		default:
			return response.Error(c, fiber.StatusInternalServerError, err.Error())
		}
	}

	return response.Success(c, fiber.StatusCreated, "user created successfully", user)
}
