package mappers

import (
	"github.com/alephshahor/Mirlo/server/models"
	"github.com/alephshahor/Mirlo/server/responses"
)

func UserModelToUserResponse(userModel models.User) responses.User {
	return responses.User{
		ID:       userModel.ID,
		UserName: userModel.UserName,
		Email:    userModel.Email,
	}
}
