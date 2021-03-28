package services

import (
	"github.com/stretchr/testify/mock"
)

type ServicesMock struct {
	mock.Mock
}

func InitializeServicesMock() *ServicesMock {
	return &ServicesMock{}
}

func (m *ServicesMock) User() IUserService {
	args := m.Called()
	return args.Get(0).(IUserService)
}

func (m *ServicesMock) SignUp() ISignUpService {
	args := m.Called()
	return args.Get(0).(ISignUpService)
}
