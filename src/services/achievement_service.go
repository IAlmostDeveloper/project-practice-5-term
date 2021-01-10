package services

import (
	"server/src/dto"
	"server/src/storage/interfaces"
)

type AchievementService struct{
	storage interfaces.StorageProvider
}

func NewAchievementService(storage interfaces.StorageProvider) *AchievementService{
	return &AchievementService{storage: storage}
}

func (s *AchievementService) GetAchievementById(id int) (*dto.Achievement, error){
	return s.storage.AchievementRepository().GetAchievementById(id)
}

func (s *AchievementService) GetAllAchievements() ([]*dto.Achievement, error){
	return s.storage.AchievementRepository().GetAllAchievements()
}

func (s *AchievementService) CompleteAchievement(id int, userId int, achieveDate *dto.TimeJson) error {
	return s.storage.AchievementRepository().CompleteAchievement(id, userId, achieveDate)
}