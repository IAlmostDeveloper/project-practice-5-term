package interfaces

import (
	"server/src/dto"
	"time"
)

type UserServiceProvider interface {
	GetUserById(id string) (*dto.User, error)
	AuthorizeUser(accessToken string) (string, error)
	AuthenticateUser(user *dto.User) (string, error)
	GenerateAndSaveToken(user *dto.User) (string, error)
	RegisterUser(user *dto.User) error
	GetAccessTokenTTL() time.Duration
	GetUserArticles(userId string) ([]*dto.Article, error)
	GetUserFocusingExercises(userId string) ([]*dto.FocusingExercise, error)
	GetUserMeditationExercises(userId string) ([]*dto.MeditationExercise, error)
	GetUserAchievements(userId string) ([]*dto.Achievement, error)
}
