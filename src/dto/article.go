package dto

type Article struct{
	ArticleId int `json:"article_id" db:"ArticleId"`
	Name string `json:"name" db:"Name"`
	Content string `json:"content" db:"Content"`
	CreateDate *TimeJson `json:"create_date" db:"CreateDate"`
}
