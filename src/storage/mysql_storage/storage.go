package mysql_storage

import (
	"github.com/jmoiron/sqlx"
	"server/src/storage/interfaces"
)

type Storage struct {
	db             *sqlx.DB
	userRepository interfaces.UserRepositoryProvider
	exerciseRepository interfaces.ExerciseRepositoryProvider
	articleRepository interfaces.ArticleRepositoryProvider
	achievementRepository interfaces.AchievementRepositoryProvider
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

func (storage *Storage) ExerciseRepository() interfaces.ExerciseRepositoryProvider {
	if storage.exerciseRepository != nil {
		return storage.exerciseRepository
	}

	storage.exerciseRepository = &ExerciseRepository{
		db: storage.db,
	}

	return storage.exerciseRepository
}

func (storage *Storage) ArticleRepository() interfaces.ArticleRepositoryProvider {
	if storage.articleRepository != nil {
		return storage.articleRepository
	}

	storage.articleRepository = &ArticleRepository{
		db: storage.db,
	}

	return storage.articleRepository
}

func (storage *Storage) AchievementRepository() interfaces.AchievementRepositoryProvider {
	if storage.achievementRepository != nil {
		return storage.achievementRepository
	}

	storage.achievementRepository = &AchievementRepository{
		db: storage.db,
	}

	return storage.achievementRepository
}
