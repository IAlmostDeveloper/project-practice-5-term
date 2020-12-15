package interfaces

import (
	"server/src/dto"
	"time"
)

type UserServiceProvider interface {
	GetUserById(id string) (*dto.User, error)
	AuthorizeUser(accessToken string) (string, error)
	AuthenticateUser(user *dto.User) (string, error)
	RegisterUser(user *dto.User) error
	ResolveUser(user *dto.User) error
	GetAccessTokenTTL() time.Duration
}
