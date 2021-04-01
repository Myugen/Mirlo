package responses

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}
