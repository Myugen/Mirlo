package repositories

import "github.com/alephshahor/Mirlo/server/models"

type IUserRepository interface {
	Create(user *models.User) error
	FindByUserName(userName string) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	_, err := DB().Model(user).Insert()
	return err
}

func (r *userRepository) FindByUserName(userName string) (models.User, error) {
	var user models.User
	var err = DB().Model(&user).
		Where("username = ?", userName).
		Select()
	return user, err
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	var err = DB().Model(&user).
		Where("email = ?", email).
		Select()
	return user, err
}
