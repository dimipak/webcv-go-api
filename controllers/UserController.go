package controllers

import (
	"app/resources"
	res "app/responses"
	"app/services"
	"app/systemService/authentication"
	v "app/validations"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Register(w http.ResponseWriter, r *http.Request) {

	userRegisterRequest, err := v.UserRegisterValidation(r)
	if err != nil {
		res.JsonResponse(res.BadRequestResponse{
			W:       &w,
			Message: err.Error(),
		})
		return
	}

	user, err := services.UserRegister(userRegisterRequest)
	if err != nil {
		res.JsonResponse(res.BadRequestResponse{
			W:       &w,
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(res.SuccessResponse{
		W:       &w,
		Message: "REGISTERED",
		Data:    resources.UserResource(user),
	})
}

func ActivateUser(w http.ResponseWriter, r *http.Request) {

	activatedKey := mux.Vars(r)["key"]

	user, err := services.ActivateUser(activatedKey)
	if err != nil {
		res.JsonResponse(res.BadRequestResponse{
			W:       &w,
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(res.SuccessResponse{
		W:       &w,
		Message: "USER_ACTIVATED",
		Data:    resources.UserResource(user),
	})
}

func Login(w http.ResponseWriter, r *http.Request) {

	UserLoginRequest, err := v.UserLoginValidation(r)
	if err != nil {
		res.JsonResponse(res.BadRequestResponse{
			W:       &w,
			Message: err.Error(),
		})
		return
	}

	user, err := services.Login(UserLoginRequest)
	if err != nil {
		res.JsonResponse(res.BadRequestResponse{
			W:       &w,
			Message: err.Error(),
		})
		return
	}

	tokenString, err := authentication.Sign(user)
	if err != nil {
		res.JsonResponse(res.BadRequestResponse{
			W:       &w,
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(res.SuccessWithTokenResponse{
		W:       &w,
		Message: "LOGGED_IN",
		Token:   tokenString,
		Data:    resources.UserResource(user),
	})
}

func GetUserProfiles(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(mux.Vars(r)["user_id"])

	profiles, err := services.GetUserProfiles(userId)
	if err != nil {
		res.JsonResponse(res.BadRequestResponse{
			W:       &w,
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(res.SuccessResponse{
		W:       &w,
		Message: "USER_PROFILES",
		Data:    resources.UserProfilesResource(profiles),
	})
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	userId, _ := strconv.Atoi(vars["user_id"])

	profileId, _ := strconv.Atoi(vars["profile_id"])

	profile, err := services.GetUserProfile(userId, profileId)
	if err != nil {
		res.JsonResponse(res.BadRequestResponse{
			W:       &w,
			Message: err.Error(),
		})
	}

	res.JsonResponse(res.SuccessResponse{
		W:       &w,
		Message: "PROFILE_RETRIEVED",
		Data:    resources.ProfileResource(profile),
	})
}

func ActivateProfile(w http.ResponseWriter, r *http.Request) {

	profile, err := services.ActivateUserProfile(r)
	if err != nil {
		res.JsonResponse(res.BadRequestResponse{
			W:       &w,
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(res.SuccessResponse{
		W:       &w,
		Message: "PROFILE_ACTIVATED",
		Data:    resources.UserProfileResource(profile),
	})
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(mux.Vars(r)["user_id"])

	request, err := v.CreateProfileValidation(r)
	if err != nil {
		res.JsonResponse(res.BadRequestResponse{
			W:       &w,
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(res.SuccessResponse{
		W:       &w,
		Message: "PROFILE_CREATED",
		Data:    resources.ProfileResource(services.CreateProfile(userId, request)),
	})
}
