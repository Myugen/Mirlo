package services

import (
	"fmt"
	"os"
	"testing"

	"github.com/alephshahor/Mirlo/server/utils"

	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/requests"

	"github.com/alephshahor/Mirlo/server/repositories"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
}

func (suite *UserServiceTestSuite) SetupTest() {
	configFilePath := os.Getenv("GOPATH") + "/src/github.com/alephshahor/Mirlo/server/.env"
	if err := godotenv.Load(configFilePath); err != nil {
		panic(fmt.Errorf("Fatal error loading .env file: %s \n", err))
	}
}

func (suite *UserServiceTestSuite) TestNewUserService() {
	var userRepository repositories.UserRepository
	var userService = NewUserService(&userRepository)

	assert.NotNil(suite.T(), userService)
}

func (suite *UserServiceTestSuite) TestCreateUser() {
	var userRepository repositories.UserRepository
	var userService = NewUserService(&userRepository)

	var user models.User
	var err error

	var newUserReq = requests.NewUserRequest{
		UserName: utils.RandString(15),
		Password: utils.RandString(15),
		Email:    utils.RandString(15),
	}

	user, err = userService.Create(newUserReq)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), user)

	assert.Equal(suite.T(), newUserReq.UserName, user.UserName)
	assert.Equal(suite.T(), newUserReq.Email, user.Email)
	assert.NotEqual(suite.T(), newUserReq.Password, user.Password)

	assert.True(suite.T(), utils.PasswordMatch(newUserReq.Password, user.Password))
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
