package controllers

import (
	"app/resources"
	res "app/responses"
	"app/services"
	"app/system"

	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct{}

func (u *UserController) Register(w http.ResponseWriter, r *http.Request) {

	err := userRegisterRequest.ValidateRequest(r)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	user, err := services.UserRegister(userRegisterRequest)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "REGISTERED",
		Data:    resources.UserResource(user),
	})
}

func (u *UserController) ActivateUser(w http.ResponseWriter, r *http.Request) {

	activatedKey := mux.Vars(r)["key"]

	user, err := services.ActivateUser(activatedKey)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "USER_ACTIVATED",
		Data:    resources.UserResource(user),
	})
}

func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {

	err := userLoginRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	user, err := services.Login(userLoginRequest)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	auth := system.Authentication{
		UserId:   user.UserId,
		Username: user.Username,
		Passwords: system.Passwords{
			Password: userLoginRequest.Password,
			Hashed:   user.Password,
		},
	}

	tokenString, err := auth.Sign()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessWithTokenResponse{
		Message: "LOGGED_IN",
		Token:   tokenString,
		Data:    resources.UserResource(user),
	})
}
