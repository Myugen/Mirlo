package mappers

import (
	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/responses"
)

func UserModelToNewUserResponse(userModel models.User) responses.NewUser {
	return responses.NewUser{
		ID:       userModel.ID,
		UserName: userModel.UserName,
		Email:    userModel.Email,
	}
}
