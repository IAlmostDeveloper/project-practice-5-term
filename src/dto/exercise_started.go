package dto

type ExerciseStarted struct {
	ExerciseId int `json:"exercise_id" db:"ExerciseId"`
	UserId int `json:"user_id" db:"UserId"`
	IsCompleted bool `json:"is_completed" db:"IsCompleted"`
	StartDate *TimeJson `json:"start_date" db:"StartDate"`
	CompleteDate *TimeJson `json:"complete_date" db:"CompleteDate"`
}
