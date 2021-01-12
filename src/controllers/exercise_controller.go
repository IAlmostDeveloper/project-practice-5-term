package controllers

import "server/src/services/interfaces"

type ExerciseController struct {
	exerciseService interfaces.ExerciseServiceProvider
}

func NewExerciseController(exerciseService interfaces.ExerciseServiceProvider) *ExerciseController{
	return &ExerciseController{exerciseService: exerciseService}
}
