// +build unit_test

package services

import (
	"testing"

	"github.com/alephshahor/Mirlo/server/utils"

	"github.com/alephshahor/Mirlo/server/errors"

	"github.com/go-pg/pg/v10"

	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/requests"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type LogInServiceUnitTestSuite struct {
	suite.Suite
}

func (suite *LogInServiceUnitTestSuite) SetupTest() {}

func (suite *LogInServiceUnitTestSuite) TestLoginNonExistingUser() {
	var servicesMock = InitializeServicesMock()
	var userServiceMock = &UserServiceMock{}
	servicesMock.On("User").Return(userServiceMock)

	var logInService = NewLogInService(servicesMock)
	assert.NotNil(suite.T(), logInService)

	var userCredentials = requests.UserCredentials{
		UserName: "JoeDoe",
		Email:    "",
		Password: "1234",
	}

	userServiceMock.On("FindByUserName", userCredentials.UserName).Return(models.User{}, pg.ErrNoRows)
	userServiceMock.On("FindByEmail", userCredentials.Email).Return(models.User{}, pg.ErrNoRows)

	var err error
	var jwt string

	_, jwt, err = logInService.LogIn(userCredentials)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.ErrUserNotFound, err)
	assert.Equal(suite.T(), "", jwt)
}

func (suite *LogInServiceUnitTestSuite) TestLoginWithUserName() {
	var err error

	var servicesMock = InitializeServicesMock()
	var userServiceMock = &UserServiceMock{}
	servicesMock.On("User").Return(userServiceMock)

	var logInService = NewLogInService(servicesMock)
	assert.NotNil(suite.T(), logInService)

	var userCredentials = requests.UserCredentials{
		UserName: "JoeDoe",
		Email:    "",
		Password: "1234",
	}

	var hashedPass string
	hashedPass, err = utils.HashPassword("1234")

	assert.Nil(suite.T(), err)

	userServiceMock.On("FindByUserName", userCredentials.UserName).Return(models.User{
		UserName: "JoeDoe",
		Email:    "JoeDoe@email.com",
		Password: hashedPass,
	}, nil)

	userServiceMock.On("FindByEmail", userCredentials.Email).Return(models.User{}, pg.ErrNoRows)

	var jwt string

	_, jwt, err = logInService.LogIn(userCredentials)

	assert.Nil(suite.T(), err)
	assert.NotZero(suite.T(), len(jwt))
}

func (suite *LogInServiceUnitTestSuite) TestLoginWithEmail() {
	var err error

	var servicesMock = InitializeServicesMock()
	var userServiceMock = &UserServiceMock{}
	servicesMock.On("User").Return(userServiceMock)

	var logInService = NewLogInService(servicesMock)
	assert.NotNil(suite.T(), logInService)

	var userCredentials = requests.UserCredentials{
		UserName: "",
		Email:    "JoeDoe@email.com",
		Password: "1234",
	}

	var hashedPass string
	hashedPass, err = utils.HashPassword("1234")

	assert.Nil(suite.T(), err)

	userServiceMock.On("FindByEmail", userCredentials.Email).Return(models.User{
		UserName: "JoeDoe",
		Email:    "JoeDoe@email.com",
		Password: hashedPass,
	}, nil)

	userServiceMock.On("FindByUserName", userCredentials.UserName).Return(models.User{}, pg.ErrNoRows)

	var jwt string

	_, jwt, err = logInService.LogIn(userCredentials)

	assert.Nil(suite.T(), err)
	assert.NotZero(suite.T(), len(jwt))
}

func TestLogInServiceUnitTestSuite(t *testing.T) {
	suite.Run(t, new(LogInServiceUnitTestSuite))
}
