package mysql_storage

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"server/src/dto"
)

type ExerciseRepository struct {
	db *sqlx.DB
}

func (repo *ExerciseRepository) GetFocusingExerciseById(id int) (*dto.FocusingExercise, error) {
	selectStatement := "SELECT * FROM `FocusingExercises` WHERE ExerciseId = ?"
	exercise := &dto.FocusingExercise{}
	if err := repo.db.Get(exercise, selectStatement, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return exercise, nil
}

func (repo *ExerciseRepository) GetMeditationExerciseById(id int) (*dto.MeditationExercise, error) {
	selectStatement := "SELECT * FROM `MeditationExercises` WHERE ExerciseId = ?"
	exercise := &dto.MeditationExercise{}
	if err := repo.db.Get(exercise, selectStatement, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return exercise, nil
}

func (repo *ExerciseRepository) GetAvailableFocusingExercises() ([]*dto.FocusingExercise, error) {
	selectStatement := "SELECT ExerciseId, Name, ExerciseTime FROM FocusingExercises"
	exercises := &[]*dto.FocusingExercise{}
	err := repo.db.Select(exercises, selectStatement)
	if err != nil {
		return nil, err
	}
	return *exercises, err
}

func (repo *ExerciseRepository) GetAvailableMeditationExercises() ([]*dto.MeditationExercise, error) {
	selectStatement := "SELECT ExerciseId, Name, ExerciseTime FROM MeditationExercises"
	exercises := &[]*dto.MeditationExercise{}
	err := repo.db.Select(exercises, selectStatement)
	if err != nil {
		return nil, err
	}
	return *exercises, err
}
