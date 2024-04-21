package types

type User struct {
	Username string `query:"username"`
	Email    string `query:"email"`
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
