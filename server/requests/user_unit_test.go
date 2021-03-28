package requests

import (
	"fmt"
	"github.com/alephshahor/Mirlo/server/utils"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserRequestTestSuite struct {
	suite.Suite
}

func (suite *UserRequestTestSuite) SetupTest() {
}

func (suite *UserRequestTestSuite) TestNewUserRequestValidator() {
	var err error
	var validate = validator.New()
	var validationErrs validator.ValidationErrors
	var newUserReq NewUser
	var correctNewUserReq = NewUser{
		UserName: utils.RandString(16),
		Password: utils.RandString(16),
		Email:    fmt.Sprintf("%v@email.com", utils.RandString(16)),
	}
	assert.Nil(suite.T(), validate.Struct(correctNewUserReq))

	// Empty UserName
	newUserReq = correctNewUserReq
	newUserReq.UserName = ""
	err = validate.Struct(newUserReq)
	validationErrs = err.(validator.ValidationErrors)
	assert.Len(suite.T(), validationErrs, 1)
	assert.Equal(suite.T(), "required", validationErrs[0].Tag())

	// Long UserName
	newUserReq = correctNewUserReq
	newUserReq.UserName = utils.RandString(21)
	err = validate.Struct(newUserReq)
	validationErrs = err.(validator.ValidationErrors)
	assert.Len(suite.T(), validationErrs, 1)
	assert.Equal(suite.T(), "lte", validationErrs[0].Tag())

	// Empty Email
	newUserReq = correctNewUserReq
	newUserReq.Email = ""
	err = validate.Struct(newUserReq)
	validationErrs = err.(validator.ValidationErrors)
	assert.Len(suite.T(), validationErrs, 1)
	assert.Equal(suite.T(), "required", validationErrs[0].Tag())

	// Email
	newUserReq = correctNewUserReq
	newUserReq.Email = "111"
	err = validate.Struct(newUserReq)
	validationErrs = err.(validator.ValidationErrors)
	assert.Len(suite.T(), validationErrs, 1)
	assert.Equal(suite.T(), "email", validationErrs[0].Tag())

	// Empty Password
	newUserReq = correctNewUserReq
	newUserReq.Password = ""
	err = validate.Struct(newUserReq)
	validationErrs = err.(validator.ValidationErrors)
	assert.Len(suite.T(), validationErrs, 1)
	assert.Equal(suite.T(), "required", validationErrs[0].Tag())

	// Short Password
	newUserReq = correctNewUserReq
	newUserReq.Password = utils.RandString(15)
	err = validate.Struct(newUserReq)
	validationErrs = err.(validator.ValidationErrors)
	assert.Len(suite.T(), validationErrs, 1)
	assert.Equal(suite.T(), "gte", validationErrs[0].Tag())

	// Long Password
	newUserReq = correctNewUserReq
	newUserReq.Password = utils.RandString(129)
	err = validate.Struct(newUserReq)
	validationErrs = err.(validator.ValidationErrors)
	assert.Len(suite.T(), validationErrs, 1)
	assert.Equal(suite.T(), "lte", validationErrs[0].Tag())
}

func TestUserRequestTestSuite(t *testing.T) {
	suite.Run(t, new(UserRequestTestSuite))
}
