package repositories

type IRepositories interface {
	User() IUserRepository
}

type repositories struct {
	userRepository IUserRepository
}

var repositoriesInstance *repositories

func Repositories() *repositories {
	if repositoriesInstance == nil {
		repositoriesInstance = initializeRepositories()
	}
	return repositoriesInstance
}

func initializeRepositories() *repositories {
	return &repositories{
		userRepository: NewUserRepository(),
	}
}

func (r *repositories) User() IUserRepository {
	return r.userRepository
}
