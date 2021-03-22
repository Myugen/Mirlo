package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/alephshahor/Mirlo/server/utils"

	"github.com/alephshahor/Mirlo/server/services"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	handlers *Handlers
}

func (suite *UserHandlerTestSuite) SetupTest() {
	configFilePath := os.Getenv("GOPATH") + "/src/github.com/alephshahor/Mirlo/server/.env"
	if err := godotenv.Load(configFilePath); err != nil {
		panic(fmt.Errorf("Fatal error loading .env file: %s \n", err))
	}
	e := echo.New()
	suite.handlers = InitializeHandlers(e, services.Services())
}

func (suite *UserHandlerTestSuite) TestCreateUser() {
	e := echo.New()
	var newUserJSON = fmt.Sprintf(`{
		"username":	"%v",
		"email": "%v",
		"password": "%v"
	}`, utils.RandString(16), utils.RandString(16), utils.RandString(16))

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(newUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(suite.T(), suite.handlers.userHandler.Create(c)) {
		assert.Equal(suite.T(), http.StatusCreated, rec.Code)
	}
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
