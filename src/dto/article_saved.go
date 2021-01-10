package dto

type ArticleSaved struct {
	Id        int       `json:"id" db:"Id"`
	ArticleId int       `json:"article_id" db:"ArticleId"`
	UserId    int       `json:"user_id" db:"UserId"`
	SaveDate  *TimeJson `json:"save_date" db:"SaveDate"`
}
