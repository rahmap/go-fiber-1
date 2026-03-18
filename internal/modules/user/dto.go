package user

type CreateRequest struct {
	Name    string `json:"name" example:"Rahma AP"`
	Age     int    `json:"age" example:"24"`
	Address string `json:"address" example:"Yogyakarta"`
}

type Response struct {
	ID      uint   `json:"id" example:"1"`
	Name    string `json:"name" example:"Rahma AP"`
	Age     int    `json:"age" example:"24"`
	Address string `json:"address" example:"Yogyakarta"`
}

type ItemEnvelope struct {
	Status  bool     `json:"status" example:"true"`
	Message string   `json:"message" example:"success"`
	Data    Response `json:"data"`
}

type ListEnvelope struct {
	Status  bool       `json:"status" example:"true"`
	Message string     `json:"message" example:"success"`
	Data    []Response `json:"data"`
}

func toResponse(user User) Response {
	return Response{
		ID:      user.ID,
		Name:    user.Name,
		Age:     user.Age,
		Address: user.Address,
	}
}

func toResponses(users []User) []Response {
	responses := make([]Response, 0, len(users))
	for _, user := range users {
		responses = append(responses, toResponse(user))
	}

	return responses
}
