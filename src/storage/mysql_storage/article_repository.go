package mysql_storage

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"server/src/dto"
)

type ArticleRepository struct {
	db *sqlx.DB
}

func (repo *ArticleRepository) GetArticleById(id int) (*dto.Article, error) {
	selectStatement := "SELECT * FROM `Articles` WHERE ArticleId = ?"
	article := &dto.Article{}
	if err := repo.db.Get(article, selectStatement, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return article, nil
}

func (repo *ArticleRepository) GetAvailableArticles() ([]*dto.Article, error) {
	selectStatement := "SELECT ArticleId, Name, Content, CreateDate FROM Articles"
	articles := &[]*dto.Article{}
	err := repo.db.Select(articles, selectStatement)
	if err != nil {
		return nil, err
	}
	return *articles, err
}
