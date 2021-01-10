package interfaces

import "server/src/dto"

type AchievementServiceProvider interface {
	GetAchievementById(id int) (*dto.Achievement, error)
	GetAllAchievements() ([]*dto.Achievement, error)
	CompleteAchievement(id int, userId int, achieveDate *dto.TimeJson) error
}
