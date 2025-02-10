package structs

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
