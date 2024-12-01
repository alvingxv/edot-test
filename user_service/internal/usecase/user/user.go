package useruc

import (
	"user-service/internal/interfaces/repository"
	"user-service/internal/interfaces/usecase"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) usecase.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}
