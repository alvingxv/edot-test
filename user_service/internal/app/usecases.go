package app

import (
	"user-service/internal/interfaces/usecase"
	useruc "user-service/internal/usecase/user"
)

type Usecases struct {
	UserUsecase usecase.UserUsecase
}

func NewUsecases(repos *Repositories) *Usecases {

	return &Usecases{
		UserUsecase: useruc.NewUserUsecase(repos.userRepository),
	}
}
