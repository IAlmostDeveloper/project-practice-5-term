package mysql_storage

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"server/src/dto"
)

type AchievementRepository struct {
	db *sqlx.DB
}

func (repo *AchievementRepository) GetAchievementById(id int) (*dto.Achievement, error) {
	selectStatement := "SELECT * FROM `Achievements` WHERE AchievementId = ?"
	achievement := &dto.Achievement{}
	if err := repo.db.Get(achievement, selectStatement, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return achievement, nil
}

func (repo *AchievementRepository) GetAllAchievements() ([]*dto.Achievement, error) {
	selectStatement := "SELECT AchievementId, Name, Description FROM Achievements"
	achievements := &[]*dto.Achievement{}
	err := repo.db.Select(achievements, selectStatement)
	if err != nil {
		return nil, err
	}
	return *achievements, err
}

func (repo *AchievementRepository) CompleteAchievement(id int, userId int, achieveDate *dto.TimeJson) error {
	insertStatement := "INSERT INTO `AchievementsAchieved` (AchievementId, UserId, AchieveDate) VALUES (?, ?, ?)"
	if _, err := repo.db.Exec(insertStatement, id, userId, achieveDate); err != nil {
		return err
	}
	return nil
}
