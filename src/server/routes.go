package server

import (
"fmt"
"github.com/gorilla/handlers"
"net/http"
)

func (s *server) ConfigureRouter() {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	s.router.Use(handlers.CORS(headers, methods, origins))

	getRouter := s.router.Methods(http.MethodGet, http.MethodOptions).Subrouter()
	postRouter := s.router.Methods(http.MethodPost, http.MethodOptions).Subrouter()
	//deleteRouter := s.router.Methods(http.MethodDelete, http.MethodOptions).Subrouter()
	//putRouter := s.router.Methods(http.MethodPut, http.MethodOptions).Subrouter()

	// Example
	//getRouter.HandleFunc("/", HelloWorld)
	//getRouter.HandleFunc("/task/{id:[0-9]+}", s.authController.AuthorizationMW(s.authController.AuthorizationMW(s.taskController.GetTaskByID)))
	//getRouter.HandleFunc("/task", s.authController.AuthorizationMW(s.taskController.GetTasks))
	//getRouter.HandleFunc("/whoami", s.authController.AuthorizationMW(s.authController.GetUser))
	//getRouter.HandleFunc("/google-auth", s.authController.GoogleLogin)
	//getRouter.HandleFunc("/google-callback", s.authController.GoogleCallback)
	//getRouter.HandleFunc("/feedback", s.feedbackController.GetAllFeedback)
	//
	//postRouter.HandleFunc("/task", s.authController.AuthorizationMW(s.taskController.CreateTask))
	//postRouter.HandleFunc("/tag", s.authController.AuthorizationMW(s.tagController.AddToTask))
	//postRouter.HandleFunc("/feedback", s.feedbackController.AddFeedback)
	//
	//deleteRouter.HandleFunc("/task/{id:[0-9]+}", s.authController.AuthorizationMW(s.taskController.RemoveTaskByID))
	//deleteRouter.HandleFunc("/tag", s.authController.AuthorizationMW(s.tagController.RemoveFromTask))
	//
	//putRouter.HandleFunc("/task/{id:[0-9]+}", s.authController.AuthorizationMW(s.taskController.UpdateTask))
	// End of Example

	// Эндпоинты для аутентификации и регистрации
	getRouter.HandleFunc("/", HelloWorld)
	getRouter.HandleFunc("/google-auth", s.userController.AuthenticateWithGoogle)
	getRouter.HandleFunc("/google-callback", s.userController.GoogleCallback)
	postRouter.HandleFunc("/register", s.userController.Register)
	postRouter.HandleFunc("/authenticate", s.userController.Authenticate)

	// Эндпоинты для данных пользователя
	getRouter.HandleFunc("/profile", s.userController.AuthorizationMW(s.userController.GetUserProfile))
	getRouter.HandleFunc("/user/achievements", s.userController.AuthorizationMW(s.userController.GetUserAchievements))
	getRouter.HandleFunc("/user/articles", s.userController.AuthorizationMW(nil))
	// Эндпоинты для медитаций
	getRouter.HandleFunc("/meditation", s.userController.AuthorizationMW(nil))
	// Эндпоинты для фокусировок
	getRouter.HandleFunc("/focusing", s.userController.AuthorizationMW(nil))
	// Эндпоинты для ачивок
	getRouter.HandleFunc("/achievements", s.userController.AuthorizationMW(s.achievementController.GetAllAchievements))
	postRouter.HandleFunc("/achievement/{id:[0-9]+}", s.userController.AuthorizationMW(s.achievementController.CompleteAchievement))


}

func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	var html = `<html><body><a href="/google-auth">Hello world!</a></body></html>`
	fmt.Fprint(writer, html)
}
