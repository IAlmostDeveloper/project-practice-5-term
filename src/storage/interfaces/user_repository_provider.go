package interfaces

import "server/src/dto"

type UserRepositoryProvider interface {
	Create(user *dto.User) error
	GetById(id string) (*dto.User, error)
	GetByLogin(login string) (*dto.User, error)
	GetByEmail(email string) (*dto.User, error)
	GetByLoginAndHashedPassword(login string, hashedPassword string) (*dto.User, error)
	Update(user *dto.User) error
	RemoveById(id string) error
	GetUserArticles(userId string) ([]*dto.Article, error)
	GetUserFocusingExercises(userId string) ([]*dto.FocusingExercise, error)
	GetUserMeditationExercises(userId string) ([]*dto.MeditationExercise, error)
	GetUserAchievements(userId string) ([]*dto.Achievement, error)
}
