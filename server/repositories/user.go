package repositories

import "github.com/alephshahor/Mirlo/server/models"

type IUserRepository interface {
	Create(user *models.User) error
}

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	_, err := DB().Model(user).Insert()
	return err
}
