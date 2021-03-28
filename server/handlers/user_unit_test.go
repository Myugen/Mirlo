package handlers

import (
	"fmt"
	"github.com/alephshahor/Mirlo/server/errors"
	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/services"
	"github.com/alephshahor/Mirlo/server/utils"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type UserHandlerUnitTestSuite struct {
	suite.Suite
}

func (suite *UserHandlerUnitTestSuite) TestCreateUser() {
	e := echo.New()
	var servicesMock = services.InitializeServicesMock()
	var signUpServiceMock = services.SignUpMock{}
	signUpServiceMock.On("Register", mock.Anything).Return(models.User{
		UserName: utils.RandString(16),
		Email:    utils.RandString(16),
		Password: utils.RandString(16),
	}, nil)
	servicesMock.On("SignUp").Return(&signUpServiceMock)
	var userHandler = InitializeUserHandler(e, servicesMock)

	var newUserJSON = fmt.Sprintf(`{
		"username":	"%v",
		"email": "%v@email.com",
		"password": "%v"
	}`, utils.RandString(16), utils.RandString(16), utils.RandString(16))

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(newUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(suite.T(), userHandler.Register(c)) {
		assert.Equal(suite.T(), http.StatusCreated, rec.Code)
	}
}

func (suite *UserHandlerUnitTestSuite) TestCreateWithAlreadyExistingEmail() {
	e := echo.New()
	var servicesMock = services.InitializeServicesMock()
	var signUpServiceMock = services.SignUpMock{}
	signUpServiceMock.On("Register", mock.Anything).Return(models.User{}, errors.ErrEmailAlreadyRegistered)
	servicesMock.On("SignUp").Return(&signUpServiceMock)
	var userHandler = InitializeUserHandler(e, servicesMock)

	var newUserJSON = fmt.Sprintf(`{
		"username":	"%v",
		"email": "%v@email.com",
		"password": "%v"
	}`, utils.RandString(16), utils.RandString(16), utils.RandString(16))

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(newUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var err = userHandler.Register(c)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.ErrEmailAlreadyRegistered, err)
}

func (suite *UserHandlerUnitTestSuite) TestCreateWithAlreadyExistingUserName() {
	e := echo.New()
	var servicesMock = services.InitializeServicesMock()
	var signUpServiceMock = services.SignUpMock{}
	signUpServiceMock.On("Register", mock.Anything).Return(models.User{}, errors.ErrUserNameAlreadyRegistered)
	servicesMock.On("SignUp").Return(&signUpServiceMock)
	var userHandler = InitializeUserHandler(e, servicesMock)

	var newUserJSON = fmt.Sprintf(`{
		"username":	"%v",
		"email": "%v@email.com",
		"password": "%v"
	}`, utils.RandString(16), utils.RandString(16), utils.RandString(16))

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(newUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var err = userHandler.Register(c)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.ErrUserNameAlreadyRegistered, err)
}

func TestUserHandlerUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerUnitTestSuite))
}
