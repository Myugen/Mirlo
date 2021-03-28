package services

import (
	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/requests"
	"github.com/stretchr/testify/mock"
)

type SignUpMock struct {
	mock.Mock
}

func (m *SignUpMock) Register(req requests.NewUser) (models.User, error) {
	args := m.Called(req)
	return args.Get(0).(models.User), args.Error(1)
}
