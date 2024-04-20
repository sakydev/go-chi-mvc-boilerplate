package types

type User struct {
	Name  string `query:"name"`
	Email string `query:"email"`
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
