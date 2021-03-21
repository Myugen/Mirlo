package models

type User struct {
	ID       uint   `pg:"id"`
	UserName string `pg:"username"`
	Email    string `pg:"email"`
	Password string `pg:"password"`
}

func NewUser(userName string, email string, password string) User {
	return User{
		UserName: userName,
		Email:    email,
		Password: password,
	}
}
