package requests

type NewUser struct {
	UserName string `json:"username" validate:"required,gt=0,lte=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=16,lte=128"`
}

type UserCredentials struct {
	UserName string `json:"username" validate:"gt=0,lte=20"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,gte=16,lte=128"`
}
