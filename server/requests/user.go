package requests

type NewUser struct {
	UserName string `json:"username" validate:"required,gt=0,lte=20"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email,gt=0,lte=320"`
}
