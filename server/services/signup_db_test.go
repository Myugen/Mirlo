//+build db_test

package services

import (
	"fmt"
	"github.com/alephshahor/Mirlo/server/errors"
	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/requests"
	"github.com/alephshahor/Mirlo/server/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type SignUpServiceDBTestSuite struct {
	suite.Suite
}

func (suite *SignUpServiceDBTestSuite) SetupTest() {
	configFilePath := os.Getenv("GOPATH") + "/src/github.com/alephshahor/Mirlo/server/.env.test"
	if err := godotenv.Load(configFilePath); err != nil {
		panic(fmt.Errorf("Fatal error loading .env file: %s \n", err))
	}
}

func (suite *SignUpServiceDBTestSuite) TestRegisterUser() {
	var services = Services()
	var signUpService = services.SignUp()
	var userService = services.User()

	assert.NotNil(suite.T(), signUpService)

	var err error

	var newUserReq = requests.NewUser{
		UserName: utils.RandString(15),
		Password: utils.RandString(15),
		Email:    utils.RandString(15),
	}

	_, err = signUpService.Register(newUserReq)

	assert.Nil(suite.T(), err)

	var foundUser models.User
	foundUser, err = userService.FindByUserName(newUserReq.UserName)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), foundUser)

	assert.Equal(suite.T(), newUserReq.UserName, foundUser.UserName)
	assert.Equal(suite.T(), newUserReq.Email, foundUser.Email)

	_, err = signUpService.Register(newUserReq)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), err, errors.ErrUserNameAlreadyRegistered)

	newUserReq.UserName = utils.RandString(15)

	_, err = signUpService.Register(newUserReq)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), err, errors.ErrEmailAlreadyRegistered)
}

func TestSignUpServiceDBTestSuite(t *testing.T) {
	suite.Run(t, new(SignUpServiceDBTestSuite))
}
