package repositories

import (
	"fmt"
	"os"
	"testing"

	"github.com/alephshahor/Mirlo/server/utils"

	"github.com/alephshahor/Mirlo/server/models"
	"github.com/stretchr/testify/assert"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
}

func (suite *UserTestSuite) SetupTest() {
	configFilePath := os.Getenv("GOPATH") + "/src/github.com/alephshahor/Mirlo/server/.env"
	if err := godotenv.Load(configFilePath); err != nil {
		panic(fmt.Errorf("Fatal error loading .env file: %s \n", err))
	}
}

func (suite *UserTestSuite) TestCreateUser() {
	var user = models.User{
		UserName: utils.RandString(15),
		Password: utils.RandString(15),
		Email:    utils.RandString(15),
	}
	var err = Repositories().User().Create(&user)
	assert.Equal(suite.T(), err, nil)

	var createdUser models.User
	err = DB().Model(&createdUser).
		Select()

	assert.Equal(suite.T(), user.ID, createdUser.ID)
	assert.Equal(suite.T(), user.Password, createdUser.Password)
	assert.Equal(suite.T(), user.Email, createdUser.Email)
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
