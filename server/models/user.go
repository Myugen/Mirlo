package models

type User struct {
	ID       uint   `pg:"id"`
	UserName string `pg:"username"`
	Email    string `pg:"email"`
	Password string `pg:"password"`
}
