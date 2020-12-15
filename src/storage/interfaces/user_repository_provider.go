package interfaces

import "server/src/dto"

type UserRepositoryProvider interface {
	Create(user *dto.User) error
	GetById(id string) (*dto.User, error)
	GetByLogin(login string) (*dto.User, error)
	Update(user *dto.User) error
	RemoveById(id string) error
}
