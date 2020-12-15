package services

import (
	"github.com/go-redis/redis/v8"
	"server/src/storage/interfaces"
)

type UserService struct {
	storage interfaces.StorageProvider
	redis   *redis.Client
}

func NewUserService(storage interfaces.StorageProvider, redis *redis.Client) *UserService {
	return &UserService{
		storage: storage,
		redis:   redis,
	}
}
