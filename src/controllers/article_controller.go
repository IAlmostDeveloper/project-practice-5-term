package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"server/src/services/interfaces"
	"strconv"
)

type ArticleController struct {
	articleService interfaces.ArticleServiceProvider
}

func NewArticleController(articleService interfaces.ArticleServiceProvider) *ArticleController {
	return &ArticleController{articleService: articleService}
}

func (controller *ArticleController) GetArticleById(writer http.ResponseWriter, request *http.Request) {
	articleId, _ := strconv.Atoi(mux.Vars(request)["id"])
	result, err := controller.articleService.GetArticleById(articleId)
	if err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusOK, result)
	return
}

func (controller *ArticleController) GetAvailableArticles(writer http.ResponseWriter, request *http.Request) {
	result, err := controller.articleService.GetAvailableArticles()
	if err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusOK, result)
	return
}

func (controller *ArticleController) GetArticlesForUser(writer http.ResponseWriter, request *http.Request) {
	result, err := controller.articleService.GetArticlesForUser(request.Context().Value(contextKeyId).(int))
	if err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusOK, result)
	return
}
