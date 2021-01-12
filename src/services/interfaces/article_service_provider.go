package interfaces

import "server/src/dto"

type ArticleServiceProvider interface {
	GetArticleById(id int) (*dto.Article, error)
	GetAvailableArticles() ([]*dto.Article, error)
	GetArticlesForUser(userId int) ([]*dto.Article, error)
}
