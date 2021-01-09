package interfaces

import "server/src/dto"

type ExerciseServiceProvider interface {
	GetFocusingExerciseById(id int) (*dto.FocusingExercise, error)
	GetMeditationExerciseById(id int) (*dto.MeditationExercise, error)
	GetAvailableFocusingExercises() ([]*dto.FocusingExercise, error)
	GetAvailableMeditationExercises() ([]*dto.MeditationExercise, error)
}
