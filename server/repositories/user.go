package repositories

import "github.com/alephshahor/Mirlo/server/models"

type UserRepository struct{}

func (r *UserRepository) Create(user *models.User) error {
	_, err := DB().Model(user).Insert()
	return err
}
