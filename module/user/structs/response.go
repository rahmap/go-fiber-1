package structs

import (
	"fiber-rest-api/module/user/model"
)

type UserResponse struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    UserDataResponse `json:"data"`
}

type UserDataResponse struct {
	ID      *int   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type UsersResponse struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    []model.User `json:"data"`
}
