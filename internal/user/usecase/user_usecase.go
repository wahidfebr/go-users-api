package usecase

import (
	"github.com/wahidfebr/go-users-api/internal/user/model"
	"github.com/wahidfebr/go-users-api/internal/user/repository"
)

type UserUseCase interface {
	CreateUser(user model.User) (model.User, error)
	GetUserByID(id string) (model.User, error)
	GetAllUsers() ([]model.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (uc *userUseCase) CreateUser(user model.User) (model.User, error) {
	return uc.repo.Create(user)
}

func (uc *userUseCase) GetUserByID(id string) (model.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *userUseCase) GetAllUsers() ([]model.User, error) {
	return uc.repo.GetAll()
}
