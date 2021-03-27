package requests

import (
	"fmt"
	"github.com/alephshahor/Mirlo/server/utils"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type UserRequestTestSuite struct {
	suite.Suite
}

func (suite *UserRequestTestSuite) SetupTest() {
	configFilePath := os.Getenv("GOPATH") + "/src/github.com/alephshahor/Mirlo/server/.env"
	if err := godotenv.Load(configFilePath); err != nil {
		panic(fmt.Errorf("Fatal error loading .env file: %s \n", err))
	}
}

func (suite *UserRequestTestSuite) TestNewUserRequestValidator() {
	var validate = validator.New()
	var correctNewUserReq = &NewUser{
		UserName: utils.RandString(16),
		Password: utils.RandString(16),
		Email:    fmt.Sprintf("%v@email.com", utils.RandString(16)),
	}
	assert.Nil(suite.T(), validate.Struct(correctNewUserReq))

	var newUserReq = correctNewUserReq
	newUserReq.UserName = ""
	assert.NotNil(suite.T(), validate.Struct(newUserReq))

	newUserReq = correctNewUserReq
	newUserReq.Email = ""
	assert.NotNil(suite.T(), validate.Struct(newUserReq))

	newUserReq = correctNewUserReq
	newUserReq.Email = utils.RandString(16)
	assert.NotNil(suite.T(), validate.Struct(newUserReq))

	newUserReq = correctNewUserReq
	newUserReq.Password = ""
	assert.NotNil(suite.T(), validate.Struct(newUserReq))
}

func TestUserRequestTestSuite(t *testing.T) {
	suite.Run(t, new(UserRequestTestSuite))
}
