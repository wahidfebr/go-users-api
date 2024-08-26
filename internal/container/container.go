package container

import (
	"github.com/wahidfebr/go-users-api/internal/user/repository"
	"github.com/wahidfebr/go-users-api/internal/user/usecase"
)

type Container struct {
	UserUseCase usecase.UserUseCase
}

func NewContainer() *Container {
	userRepo := repository.NewUserRepository()
	userUseCase := usecase.NewUserUseCase(userRepo)

	return &Container{
		UserUseCase: userUseCase,
	}
}
