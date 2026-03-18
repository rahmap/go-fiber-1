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

// GetAll godoc
// @Summary Get all users
// @Description Returns all users from the database
// @Tags Users
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} ListEnvelope
// @Failure 500 {object} response.ErrorEnvelope
// @Router /api/v1/users [get]
func (h *Handler) GetAll(c *fiber.Ctx) error {
	users, err := h.service.GetAll()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "success", users)
}

// GetByID godoc
// @Summary Get user by ID
// @Description Returns one user by path parameter
// @Tags Users
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} ItemEnvelope
// @Failure 400 {object} response.ErrorEnvelope
// @Failure 404 {object} response.ErrorEnvelope
// @Failure 500 {object} response.ErrorEnvelope
// @Router /api/v1/users/{id} [get]
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

// Create godoc
// @Summary Create a user
// @Description Creates a new user in the database
// @Tags Users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body CreateRequest true "Create user payload"
// @Success 201 {object} ItemEnvelope
// @Failure 400 {object} response.ErrorEnvelope
// @Failure 500 {object} response.ErrorEnvelope
// @Router /api/v1/users [post]
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
