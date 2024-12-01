package app

import (
	"user-service/internal/interfaces/repository"
	userrepo "user-service/internal/repository/user"
)

type Repositories struct {
	userRepository repository.UserRepository
}

func NewRepos(dependencies *Dependencies) *Repositories {
	return &Repositories{
		userRepository: userrepo.NewUserRepository(dependencies.sqlitedb),
	}
}
