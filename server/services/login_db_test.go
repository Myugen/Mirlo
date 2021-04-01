//+build db_test

package services

import (
	"fmt"
	"os"
	"testing"

	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/requests"
	"github.com/alephshahor/Mirlo/server/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoginServiceDBTestSuite struct {
	suite.Suite
}

func (suite *LoginServiceDBTestSuite) SetupTest() {
	configFilePath := os.Getenv("GOPATH") + "/src/github.com/alephshahor/Mirlo/server/.env.test"
	if err := godotenv.Load(configFilePath); err != nil {
		panic(fmt.Errorf("Fatal error loading .env file: %s \n", err))
	}
}

func (suite *LoginServiceDBTestSuite) TestLoginUser() {
	var err error

	var services = Services()
	var signUpService = services.SignUp()

	var newUserReq = requests.NewUser{
		UserName: utils.RandString(15),
		Password: utils.RandString(15),
		Email:    utils.RandString(15),
	}

	_, err = signUpService.Register(newUserReq)
	assert.Nil(suite.T(), err)

	var loginService = services.LogIn()

	var jwtToken string
	var user models.User

	var userCredentials = requests.UserCredentials{
		UserName: newUserReq.UserName,
		Password: newUserReq.Password,
	}

	user, jwtToken, err = loginService.LogIn(userCredentials)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), user)
	assert.NotZero(suite.T(), user.ID)
	assert.Equal(suite.T(), newUserReq.UserName, user.UserName)
	assert.Equal(suite.T(), newUserReq.Email, user.Email)
	assert.True(suite.T(), utils.PasswordMatch(newUserReq.Password, user.Password))
	assert.NotZero(suite.T(), len(jwtToken))
}

func TestLoginServiceDBTestSuite(t *testing.T) {
	suite.Run(t, new(LoginServiceDBTestSuite))
}
