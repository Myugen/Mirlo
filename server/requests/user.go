package requests

type NewUserRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
