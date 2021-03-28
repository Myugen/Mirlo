package repositories

import (
	"github.com/alephshahor/Mirlo/server/models"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) FindByUserName(userName string) (models.User, error) {
	args := m.Called(userName)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *UserRepositoryMock) FindByEmail(email string) (models.User, error) {
	args := m.Called(email)
	return args.Get(0).(models.User), args.Error(1)
}
