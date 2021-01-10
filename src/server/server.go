package server

import (
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"net/http"
	"server/src/controllers"
	"server/src/services"
	"server/src/storage/interfaces"
)

type server struct {
	router             *mux.Router
	storage            interfaces.StorageProvider
	userController     *controllers.UserController
	exerciseController *controllers.ExerciseController
	articleController *controllers.ArticleController
	achievementController *controllers.AchievementController
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(storage interfaces.StorageProvider, redis *redis.Client) *server {
	server := &server{
		router:  mux.NewRouter(),
		storage: storage,
		userController: controllers.NewUserController(
			services.NewUserService(storage, redis), services.NewPasswordService(),
		),
		exerciseController: controllers.NewExerciseController(services.NewExerciseService(storage)),
		articleController: controllers.NewArticleController(services.NewArticleService(storage)),
		achievementController: controllers.NewAchievementController(services.NewAchievementService(storage)),
	}

	server.ConfigureRouter()
	return server
}
