package repository

import (
	"errors"
	"sync"

	"github.com/wahidfebr/go-users-api/internal/user/model"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	Create(user model.User) (model.User, error)
	GetByID(id string) (model.User, error)
	GetAll() ([]model.User, error)
}

type inMemoryUserRepository struct {
	users map[string]model.User
	mu    sync.RWMutex
}

func NewUserRepository() UserRepository {
	return &inMemoryUserRepository{
		users: make(map[string]model.User),
	}
}

func (r *inMemoryUserRepository) Create(user model.User) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return user, nil
}

func (r *inMemoryUserRepository) GetByID(id string) (model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, exists := r.users[id]
	if !exists {
		return model.User{}, ErrUserNotFound
	}
	return user, nil
}

func (r *inMemoryUserRepository) GetAll() ([]model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var users []model.User
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}
