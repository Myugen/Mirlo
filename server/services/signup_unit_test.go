package services

import (
	"github.com/alephshahor/Mirlo/server/errors"
	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/repositories"
	"github.com/alephshahor/Mirlo/server/requests"
	"github.com/alephshahor/Mirlo/server/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ServiceUpServiceTestSuite struct {
	suite.Suite
}

func (suite *ServiceUpServiceTestSuite) SetupTest() {
}

func (suite *ServiceUpServiceTestSuite) TestRegisterNewUser() {
	var servicesMock = InitializeServicesMock()
	var repositoriesMock = repositories.InitializeRepositoriesMock()
	var userServiceMock = &UserServiceMock{}
	servicesMock.On("User").Return(userServiceMock)
	var signUpService = NewSignUpService(repositoriesMock, servicesMock)

	assert.NotNil(suite.T(), signUpService)

	var err error

	var newUserReq = requests.NewUser{
		UserName: utils.RandString(15),
		Password: utils.RandString(15),
		Email:    utils.RandString(15),
	}

	userServiceMock.On("FindByUserName", newUserReq.UserName).Return(models.User{
		UserName: "",
	}, nil)

	userServiceMock.On("FindByEmail", newUserReq.Email).Return(models.User{
		Email: "",
	}, nil)

	userServiceMock.On("Create", newUserReq).Return(models.User{}, nil)

	_, err = signUpService.Register(newUserReq)
	assert.Nil(suite.T(), err)
}

func (suite *ServiceUpServiceTestSuite) TestAlreadyRegisteredUserName() {
	var servicesMock = InitializeServicesMock()
	var repositoriesMock = repositories.InitializeRepositoriesMock()
	var userServiceMock = &UserServiceMock{}
	servicesMock.On("User").Return(userServiceMock)
	var signUpService = NewSignUpService(repositoriesMock, servicesMock)

	assert.NotNil(suite.T(), signUpService)

	var err error

	var newUserReq = requests.NewUser{
		UserName: utils.RandString(15),
		Password: utils.RandString(15),
		Email:    utils.RandString(15),
	}

	userServiceMock.On("FindByUserName", newUserReq.UserName).Return(models.User{
		UserName: newUserReq.UserName,
	}, nil)

	_, err = signUpService.Register(newUserReq)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.ErrUserNameAlreadyRegistered, err)
}

func (suite *ServiceUpServiceTestSuite) TestAlreadyRegisteredEmail() {
	var servicesMock = InitializeServicesMock()
	var repositoriesMock = repositories.InitializeRepositoriesMock()
	var userServiceMock = &UserServiceMock{}
	servicesMock.On("User").Return(userServiceMock)
	var signUpService = NewSignUpService(repositoriesMock, servicesMock)

	assert.NotNil(suite.T(), signUpService)

	var err error

	var newUserReq = requests.NewUser{
		UserName: utils.RandString(15),
		Password: utils.RandString(15),
		Email:    utils.RandString(15),
	}

	userServiceMock.On("FindByUserName", newUserReq.UserName).Return(models.User{
		UserName: "",
	}, nil)

	userServiceMock.On("FindByEmail", newUserReq.Email).Return(models.User{
		Email: utils.RandString(16),
	}, nil)

	_, err = signUpService.Register(newUserReq)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.ErrEmailAlreadyRegistered, err)
}

func TestServiceUpServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceUpServiceTestSuite))
}
