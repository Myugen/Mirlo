package services

import "github.com/alephshahor/Mirlo/server/repositories"

type IServices interface {
	User() IUserService
	SignUp() ISignUpService
	LogIn() ILogInService
}

type services struct {
	userService   IUserService
	signUpService ISignUpService
	logInService  ILogInService
}

var servicesInstance *services

func Services() IServices {
	if servicesInstance == nil {
		servicesInstance = &services{}
	}
	return servicesInstance
}

func (s *services) User() IUserService {
	if s.userService == nil {
		s.userService = NewUserService(repositories.Repositories())
	}
	return s.userService
}

func (s *services) SignUp() ISignUpService {
	if s.signUpService == nil {
		s.signUpService = NewSignUpService(s)
	}
	return s.signUpService
}

func (s *services) LogIn() ILogInService {
	if s.logInService == nil {
		s.logInService = NewLogInService(s)
	}
	return s.logInService
}
