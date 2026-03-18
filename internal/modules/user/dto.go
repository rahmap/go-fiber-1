package user

type CreateRequest struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type Response struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
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
