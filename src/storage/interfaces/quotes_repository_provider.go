package interfaces

import "server/src/dto"

type QuotesRepositoryProvider interface {
	GetQuoteById(id int) (*dto.Quote, error)
	GetAllQuotes() ([]*dto.Quote, error)
	GetUserSavedQuotes(userId int) ([]*dto.Quote, error)
	InsertQuote(quote dto.Quote) error
}
