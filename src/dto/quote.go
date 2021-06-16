package dto

type Quote struct {
	Id int `json:"id" db:"Id"`
	AuthorName string `json:"author_name" db:"AuthorName"`
	Theme string `json:"theme" db:"Theme"`
	Quote string `json:"quote" db:"Quote"`
}

type QuoteSaved struct {
	Id int `json:"id" db:"Id"`
	UserId int `json:"user_id" db:"UserId"`
	QuoteId int `json:"quote_id" db:"QuoteId"`
}