package services

import (
	"server/src/dto"
	"server/src/storage/interfaces"
)

type ExerciseService struct {
	storage interfaces.StorageProvider
}

func NewExerciseService(storage interfaces.StorageProvider) *ExerciseService {
	return &ExerciseService{storage: storage}
}

func (s *ExerciseService) GetFocusingExerciseById(id int) (*dto.FocusingExercise, error) {
	return s.storage.ExerciseRepository().GetFocusingExerciseById(id);
}

func (s *ExerciseService) GetMeditationExerciseById(id int) (*dto.MeditationExercise, error) {
	return s.storage.ExerciseRepository().GetMeditationExerciseById(id)
}

func (s *ExerciseService) GetAvailableFocusingExercises() ([]*dto.FocusingExercise, error) {
	return s.storage.ExerciseRepository().GetAvailableFocusingExercises()
}

func (s *ExerciseService) GetAvailableMeditationExercises() ([]*dto.MeditationExercise, error) {
	return s.storage.ExerciseRepository().GetAvailableMeditationExercises()
}
