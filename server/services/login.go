package services

import (
	"os"
	"time"

	"github.com/alephshahor/Mirlo/server/utils"

	"github.com/alephshahor/Mirlo/server/errors"

	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/requests"
	"github.com/dgrijalva/jwt-go"
)

type ILogInServiceServices interface {
	User() IUserService
}

type ILogInService interface {
	LogIn(userCredentials requests.UserCredentials) (models.User, string, error)
}

type logInService struct {
	services IServices
}

func NewLogInService(services IServices) *logInService {
	return &logInService{
		services: services,
	}
}

func (s *logInService) LogIn(userCredentials requests.UserCredentials) (models.User, string, error) {
	var err error
	var user models.User
	var jwtToken string

	switch {
	case userCredentials.UserName != "":
		if user, err = s.services.User().FindByUserName(userCredentials.UserName); err != nil {
			return user, "", errors.ErrUserNotFound
		}
	case userCredentials.Email != "":
		if user, err = s.services.User().FindByEmail(userCredentials.Email); err != nil {
			return user, "", errors.ErrUserNotFound
		}
	default:
		return user, "", errors.ErrInvalidLoginCredentials
	}

	if !utils.PasswordMatch(userCredentials.Password, user.Password) {
		return user, "", errors.ErrInvalidLoginCredentials
	}

	if jwtToken, err = s.generateJWT(user); err != nil {
		return user, "", err
	}

	return user, jwtToken, nil
}

func (s *logInService) generateJWT(user models.User) (string, error) {
	var err error

	var token *jwt.Token
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issuer": user.UserName,
		"exp":    time.Now().Add(time.Minute * 15).Unix(),
	})

	var tokenString string
	if tokenString, err = token.SignedString([]byte(os.Getenv("SECRET_KEY"))); err != nil {
		return "", err
	}

	return tokenString, nil
}
