package services

import (
	"server/src/dto"
	"server/src/storage/interfaces"
)

type ArticleService struct {
	storage interfaces.StorageProvider
}

func NewArticleService(storage interfaces.StorageProvider) *ArticleService {
	return &ArticleService{storage: storage}
}

func (s *ArticleService) GetArticleById(id int) (*dto.Article, error) {
	return s.storage.ArticleRepository().GetArticleById(id)
}
func (s *ArticleService) GetAvailableArticles() ([]*dto.Article, error) {
	return s.storage.ArticleRepository().GetAvailableArticles()
}

func (s *ArticleService) GetArticlesForUser(userId int) ([]*dto.Article, error) {
	return s.storage.ArticleRepository().GetArticlesForUser(userId)
}
