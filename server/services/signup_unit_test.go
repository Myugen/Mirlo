package services

import (
	"fmt"
	"github.com/alephshahor/Mirlo/server/errors"
	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/repositories"
	"github.com/alephshahor/Mirlo/server/requests"
	"github.com/alephshahor/Mirlo/server/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type MockedService struct {
	mock.Mock
	userService *MockedUserService
}

func InitializeMockedService() MockedService {
	return MockedService{
		userService: &MockedUserService{},
	}
}

func (m *MockedService) User() IUserService {
	return m.userService
}

func (m *MockedService) SignUp() ISignUpService {
	return NewSignUpService(repositories.Repositories(), m)
}

type MockedUserService struct {
	mock.Mock
}

func (m *MockedUserService) Create(req requests.NewUser) (models.User, error) {
	args := m.Called(req)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockedUserService) FindByUserName(userName string) (models.User, error) {
	args := m.Called(userName)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockedUserService) FindByEmail(email string) (models.User, error) {
	args := m.Called(email)
	return args.Get(0).(models.User), args.Error(1)
}

type ServiceUpServiceTestSuite struct {
	suite.Suite
}

func (suite *ServiceUpServiceTestSuite) SetupTest() {
	configFilePath := os.Getenv("GOPATH") + "/src/github.com/alephshahor/Mirlo/server/.env"
	if err := godotenv.Load(configFilePath); err != nil {
		panic(fmt.Errorf("Fatal error loading .env file: %s \n", err))
	}
}

func (suite *ServiceUpServiceTestSuite) TestRegisterNewUser() {
	var services = InitializeMockedService()
	var signUpService = services.SignUp()

	assert.NotNil(suite.T(), signUpService)

	var err error

	var newUserReq = requests.NewUser{
		UserName: utils.RandString(15),
		Password: utils.RandString(15),
		Email:    utils.RandString(15),
	}

	services.userService.On("FindByUserName", newUserReq.UserName).Return(models.User{
		UserName: "",
	}, nil)

	services.userService.On("FindByEmail", newUserReq.Email).Return(models.User{
		Email: "",
	}, nil)

	services.userService.On("Create", newUserReq).Return(models.User{}, nil)

	_, err = signUpService.Register(newUserReq)
	assert.Nil(suite.T(), err)
}

func (suite *ServiceUpServiceTestSuite) TestAlreadyRegisteredUserName() {
	var services = InitializeMockedService()
	var signUpService = services.SignUp()

	assert.NotNil(suite.T(), signUpService)

	var err error

	var newUserReq = requests.NewUser{
		UserName: utils.RandString(15),
		Password: utils.RandString(15),
		Email:    utils.RandString(15),
	}

	services.userService.On("FindByUserName", newUserReq.UserName).Return(models.User{
		UserName: newUserReq.UserName,
	}, nil)

	_, err = signUpService.Register(newUserReq)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.ErrUserNameAlreadyRegistered, err)
}

func (suite *ServiceUpServiceTestSuite) TestAlreadyRegisteredEmail() {
	var services = InitializeMockedService()
	var signUpService = services.SignUp()

	assert.NotNil(suite.T(), signUpService)

	var err error

	var newUserReq = requests.NewUser{
		UserName: utils.RandString(15),
		Password: utils.RandString(15),
		Email:    utils.RandString(15),
	}

	services.userService.On("FindByUserName", newUserReq.UserName).Return(models.User{
		UserName: "",
	}, nil)

	services.userService.On("FindByEmail", newUserReq.Email).Return(models.User{
		Email: utils.RandString(16),
	}, nil)

	_, err = signUpService.Register(newUserReq)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.ErrEmailAlreadyRegistered, err)
}

func TestServiceUpServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceUpServiceTestSuite))
}
