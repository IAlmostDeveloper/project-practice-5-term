package dto

type Achievement struct {
	AchievementId int    `json:"achievement_id" db:"AchievementId"`
	Name          string `json:"name" db:"Name"`
	Description   string `json:"description" db:"Name"`
}
