package mysql_storage

import (
	"github.com/jmoiron/sqlx"
	"server/src/storage/interfaces"
)

type Storage struct {
	db             *sqlx.DB
	userRepository interfaces.UserRepositoryProvider
}

func New(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (storage *Storage) UserRepository() interfaces.UserRepositoryProvider {
	if storage.userRepository != nil {
		return storage.userRepository
	}

	storage.userRepository = &UserRepository{
		db: storage.db,
	}

	return storage.userRepository
}
