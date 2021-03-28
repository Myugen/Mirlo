package repositories

import "github.com/stretchr/testify/mock"

type RepositoriesMock struct {
	mock.Mock
}

func InitializeRepositoriesMock() *RepositoriesMock {
	return &RepositoriesMock{}
}

func (m *RepositoriesMock) User() IUserRepository {
	args := m.Called()
	return args.Get(0).(IUserRepository)
}
