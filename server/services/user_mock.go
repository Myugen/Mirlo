package services

import (
	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/requests"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) Create(req requests.NewUser) (models.User, error) {
	args := m.Called(req)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *UserServiceMock) FindByUserName(userName string) (models.User, error) {
	args := m.Called(userName)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *UserServiceMock) FindByEmail(email string) (models.User, error) {
	args := m.Called(email)
	return args.Get(0).(models.User), args.Error(1)
}
