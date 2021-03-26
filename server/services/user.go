package services

import (
	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/repositories"
	"github.com/alephshahor/Mirlo/server/requests"
	"github.com/alephshahor/Mirlo/server/utils"
)

type IUserServiceRepositories interface {
	User() repositories.IUserRepository
}

type IUserService interface {
	Create(req requests.NewUser) (models.User, error)
	FindByUserName(userName string) (models.User, error)
}

type userService struct {
	repositories IUserServiceRepositories
}

func NewUserService(repositories repositories.IRepositories) *userService {
	return &userService{
		repositories: repositories,
	}
}

func (s *userService) Create(req requests.NewUser) (models.User, error) {
	var err error
	var user models.User

	var hashedPassword string
	if hashedPassword, err = utils.HashPassword(req.Password); err != nil {
		return user, err
	}

	user = models.NewUser(req.UserName, req.Email, hashedPassword)

	err = s.repositories.User().Create(&user)

	return user, err
}

func (s *userService) FindByUserName(userName string) (models.User, error) {
	var err error
	var user models.User

	if user, err = s.repositories.User().FindByUserName(userName); err != nil {
		return user, err
	}

	return user, nil
}
