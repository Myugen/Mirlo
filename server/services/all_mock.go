package services

import (
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func InitializeServiceMock() *ServiceMock {
	return &ServiceMock{}
}

func (m *ServiceMock) User() IUserService {
	args := m.Called()
	return args.Get(0).(IUserService)
}

func (m *ServiceMock) SignUp() ISignUpService {
	args := m.Called()
	return args.Get(0).(ISignUpService)
}
