package dto

type Exercise struct {
	ExerciseId int `json:"exercise_id" db:"ExerciseId"`
	Name string `json:"name" db:"Name"`
	ExerciseTime int `json:"exercise_time" db:"ExerciseTime"`
}
