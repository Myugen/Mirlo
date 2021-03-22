package services

import "github.com/alephshahor/Mirlo/server/repositories"

type IServices interface {
	User() IUserService
}

type services struct {
	userService IUserService
}

var servicesInstance *services

func Services() *services {
	if servicesInstance == nil {
		servicesInstance = initializeServices()
	}
	return servicesInstance
}

func initializeServices() *services {
	var repositories = repositories.Repositories()
	return &services{
		userService: NewUserService(repositories),
	}
}

func (s *services) User() IUserService {
	return s.userService
}
