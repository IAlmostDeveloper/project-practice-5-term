package mysql_storage

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"server/src/dto"
)

type AchievementRepository struct{
	db *sqlx.DB
}

func (repo *AchievementRepository) GetAchievementById(id int) (*dto.Achievement, error){
	selectStatement := "SELECT * FROM `Achievements` WHERE AchievementId = ?"
	exercise := &dto.Achievement{}
	if err := repo.db.Get(exercise, selectStatement, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return exercise, nil
}

func (repo *AchievementRepository) CompleteAchievement(id int, userId int, achieveDate *dto.TimeJson) error{
	insertStatement := "INSERT INTO `AchievementsAchieved` (AchievementId, UserId, AchieveDate) VALUES (?, ?, ?)"
	if _, err := repo.db.Exec(insertStatement, id, userId, achieveDate); err != nil {
		return err
	}
	return nil
}
