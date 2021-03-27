package services

import (
	"github.com/alephshahor/Mirlo/server/errors"
	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/repositories"
	"github.com/alephshahor/Mirlo/server/requests"
	"github.com/go-pg/pg/v10"
)

type ISignUpServiceRepositories interface{}

type ISignUpServiceServices interface {
	User() IUserService
}

type ISignUpService interface {
	Register(req requests.NewUser) (models.User, error)
}

type signUpService struct {
	repositories ISignUpServiceRepositories
	services     ISignUpServiceServices
}

func NewSignUpService(repositories repositories.IRepositories, services IServices) *signUpService {
	return &signUpService{
		repositories: repositories,
		services:     services,
	}
}

func (s *signUpService) Register(req requests.NewUser) (models.User, error) {
	var err error
	var user models.User

	if user, err = s.services.User().FindByUserName(req.UserName); err != nil && err != pg.ErrNoRows {
		return user, err
	}

	if user.UserName != "" {
		return user, errors.ErrUserNameAlreadyRegistered
	}

	if user, err = s.services.User().FindByEmail(req.Email); err != nil && err != pg.ErrNoRows {
		return user, err
	}

	if user.Email != "" {
		return user, errors.ErrEmailAlreadyRegistered
	}

	if user, err = s.services.User().Create(req); err != nil {
		return user, err
	}

	return user, nil
}
