package interfaces

import "server/src/dto"

type ArticleRepositoryProvider interface {
	GetArticleById(id int) (*dto.Article, error)
	GetAvailableArticles() ([]*dto.Article, error)
	GetArticlesForUser(userId int) ([]*dto.Article, error)
}
