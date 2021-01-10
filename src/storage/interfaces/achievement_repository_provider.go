package interfaces

import "server/src/dto"

type AchievementRepositoryProvider interface {
	GetAchievementById(id int) (*dto.Achievement, error)
	CompleteAchievement(id int, userId int, achieveDate *dto.TimeJson) error
}
