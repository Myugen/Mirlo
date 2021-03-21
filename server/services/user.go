package services

import (
	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/repositories"
	"github.com/alephshahor/Mirlo/server/requests"
	"github.com/alephshahor/Mirlo/server/utils"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Create(req requests.NewUserRequest) (models.User, error) {
	var err error
	var user models.User

	var hashedPassword string
	if hashedPassword, err = utils.HashPassword(req.Password); err != nil {
		return user, err
	}

	user = models.NewUser(req.UserName, req.Email, hashedPassword)

	err = s.userRepository.Create(&user)

	return user, err
}
