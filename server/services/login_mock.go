package services

import (
	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/requests"
	"github.com/stretchr/testify/mock"
)

type LoginMock struct {
	mock.Mock
}

func (m *LoginMock) LogIn(req requests.UserCredentials) (models.User, string, error) {
	args := m.Called(req)
	return args.Get(0).(models.User), args.String(1), args.Error(2)
}
