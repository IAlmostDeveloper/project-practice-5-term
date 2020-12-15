package controllers

import (
	"net/http"
	"server/src/services/interfaces"
)

type UserController struct{
	userService interfaces.UserServiceProvider
	passwordService interfaces.PasswordServiceProvider
}

func NewUserController(userService interfaces.UserServiceProvider,
	passwordService interfaces.PasswordServiceProvider) *UserController{
	return &UserController{
		userService,
		passwordService,
	}
}

func (controller *UserController) Authenticate (writer http.ResponseWriter, request *http.Request){

}

func (controller *UserController) AuthenticateWithGoogle (writer http.ResponseWriter, request *http.Request){

}

func (controller *UserController) Register (writer http.ResponseWriter, request *http.Request){

}

func (controller *UserController) Authorize (writer http.ResponseWriter, request *http.Request){

}
