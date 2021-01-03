package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"os"
	"server/src/dto"
	"server/src/services/interfaces"
	"strings"
	"time"
)

const (
	contextKeyId = "userId"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/google-callback",
		ClientID:     os.Getenv("GOOGLE_AUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_AUTH_CLIENT_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/plus.me"},
		Endpoint: google.Endpoint,
	}
	randomState = "random"
)

type UserController struct {
	userService     interfaces.UserServiceProvider
	passwordService interfaces.PasswordServiceProvider
}

func NewUserController(userService interfaces.UserServiceProvider,
	passwordService interfaces.PasswordServiceProvider) *UserController {
	return &UserController{
		userService,
		passwordService,
	}
}

func (controller *UserController) Authenticate(writer http.ResponseWriter, request *http.Request) {
	var user dto.User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		errorJsonRespond(writer, http.StatusBadRequest, errJsonDecode)
		return
	}
	// Не ввожу поле "Password", поначалу нехешированый пароль попадает в это поле
	user.HashedPassword = controller.passwordService.EncodePassword(user.HashedPassword)
	// В AuthenticateUser попадает уже хешированный пароль
	accessToken, err := controller.userService.AuthenticateUser(&user)
	if err != nil {
		if err == errInvalidUserData {
			errorJsonRespond(writer, http.StatusBadRequest, err)
			return
		}
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	http.SetCookie(writer, &http.Cookie{
		Name:       "accessToken",
		Value:      accessToken,
		Path:       "/",
		RawExpires: time.Now().Add(controller.userService.GetAccessTokenTTL()).String(),
	})

	respondJson(writer, http.StatusOK, accessToken)
}
func (controller *UserController) Register(writer http.ResponseWriter, request *http.Request) {
	var user dto.User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		errorJsonRespond(writer, http.StatusBadRequest, errJsonDecode)
		return
	}
	// Не ввожу поле "Password", поначалу нехешированый пароль попадает в это поле
	user.HashedPassword = controller.passwordService.EncodePassword(user.HashedPassword)
	// В RegisterUser попадает уже хешированный пароль
	if err := controller.userService.RegisterUser(&user); err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusCreated, user)
}

func (controller *UserController) AuthenticateWithGoogle(writer http.ResponseWriter, request *http.Request) {
	url := googleOauthConfig.AuthCodeURL(randomState)
	fmt.Println(url)
	http.Redirect(writer, request, url, http.StatusTemporaryRedirect)
}

func (controller *UserController) GoogleCallback(writer http.ResponseWriter, request *http.Request) {
	if request.FormValue("state") != randomState {
		fmt.Println("state is not valid")
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, request.FormValue("code"))
	if err != nil {
		fmt.Printf("could not get token : %s \n", err.Error())
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("could not create get request : %s \n", err.Error())
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()
	user := &dto.User{}
	googleUser := &struct {
		Email     string `json:"email"`
		FirstName string `json:"given_name"`
		LastName  string `json:"family_name"`
		Picture   string `json:"picture"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(googleUser); err != nil {
		errorJsonRespond(writer, http.StatusBadRequest, errJsonDecode)
		return
	}
	user.Email = googleUser.Email
	user.Login = strings.Split(googleUser.Email, "@")[0]
	user.FirstName = googleUser.FirstName
	user.LastName = googleUser.LastName
	user.IsRegisteredWithGoogle = true
	user.AvatarPicture = googleUser.Picture

	// create user if not exist
	if err := controller.userService.RegisterUser(user); err != nil {
		if err.Error() == "this login or email is already used" {
			token, err := controller.userService.AuthenticateUser(user)
			if err != nil {
				errorJsonRespond(writer, http.StatusBadRequest, err)
				return
			}
			http.SetCookie(writer, &http.Cookie{
				Name:       "accessToken",
				Value:      token,
				Path:       "/",
				RawExpires: time.Now().Add(controller.userService.GetAccessTokenTTL()).String(),
			})
			respondJson(writer, http.StatusOK, user)
			return
		}
		errorJsonRespond(writer, http.StatusBadRequest, err)
		return
	}

	accessToken, err := controller.userService.GenerateAndSaveToken(user)
	if err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(writer, &http.Cookie{
		Name:       "accessToken",
		Value:      accessToken,
		Path:       "/",
		RawExpires: time.Now().Add(controller.userService.GetAccessTokenTTL()).String(),
	})
	respondJson(writer, http.StatusOK, user)
	//http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}

func (controller *UserController) GetUserProfile(writer http.ResponseWriter, request *http.Request) {
	userId := request.Context().Value(contextKeyId).(string)
	result, err := controller.userService.GetUserById(userId)
	if err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}
	respondJson(writer, http.StatusOK, result)
}

func (controller *UserController) AuthorizationMW(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("accessToken")
		if err != nil {
			errorJsonRespond(w, http.StatusUnauthorized, err)
			return
		}
		userId, err := controller.userService.AuthorizeUser(tokenCookie.Value)
		if err != nil {
			errorJsonRespond(w, http.StatusUnauthorized, err)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), contextKeyId, userId)))
	})
}
