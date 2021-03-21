package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	var err error
	var hashedPassword []byte

	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func PasswordMatch(password string, hash string) bool {
	var err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
