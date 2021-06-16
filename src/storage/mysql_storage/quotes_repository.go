package mysql_storage

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"server/src/dto"
)

type QuotesRepository struct {
	db *sqlx.DB
}

func (repo *QuotesRepository) GetQuoteById(id int) (*dto.Quote, error){
	selectStatement := "SELECT * FROM `Quotes` WHERE Id = ?"
	quote := &dto.Quote{}
	if err := repo.db.Get(quote, selectStatement, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return quote, nil
}
func (repo *QuotesRepository) GetAllQuotes() ([]*dto.Quote, error){
	selectStatement := "SELECT Id, AuthorName, Theme, Quote FROM Quotes"
	quotes := &[]*dto.Quote{}
	err := repo.db.Select(quotes, selectStatement)
	if err != nil {
		return nil, err
	}
	return *quotes, err
}
func (repo *QuotesRepository) GetUserSavedQuotes(userId int) ([]*dto.Quote, error){
		return nil, nil
}
func (repo *QuotesRepository) InsertQuote(quote dto.Quote) error{
	return nil
}
