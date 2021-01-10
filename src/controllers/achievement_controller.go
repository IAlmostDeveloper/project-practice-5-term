package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"server/src/dto"
	"server/src/services/interfaces"
	"strconv"
	"time"
)

type AchievementController struct {
	achievementService interfaces.AchievementServiceProvider
}

func NewAchievementController(achievementService interfaces.AchievementServiceProvider) *AchievementController {
	return &AchievementController{
		achievementService: achievementService,
	}
}

func (controller *AchievementController) GetAllAchievements(writer http.ResponseWriter, request *http.Request) {
	result, err := controller.achievementService.GetAllAchievements()
	if err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, errNotFound)
	}
	respondJson(writer, http.StatusOK, result)
	return
}

func (controller *AchievementController) CompleteAchievement(writer http.ResponseWriter, request *http.Request) {
	userId, _ := strconv.Atoi(request.Context().Value(contextKeyId).(string))
	achievementId, _ := strconv.Atoi(mux.Vars(request)["id"])
	dateNow := new(dto.TimeJson)
	if err := dateNow.UnmarshalJSON([]byte(time.Now().Format(dto.DateFormat))); err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, errNotFound)
		return
	}
	if err := controller.achievementService.CompleteAchievement(achievementId, userId, dateNow); err != nil{
		errorJsonRespond(writer, http.StatusInternalServerError, errNotFound)
		return
	}

}
