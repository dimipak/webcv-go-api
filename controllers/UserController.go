package controllers

import (
	"app/resources"
	rw "app/responses"
	"app/services"
	"app/systemService/authentication"
	v "app/validations"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Register(w http.ResponseWriter, r *http.Request) {

	res := rw.ResponseWriter{W: &w}

	userRegisterRequest, err := v.UserRegisterValidation(r)
	if err != nil {
		res.BadRequest(err.Error())
		return
	}

	user, err := services.UserRegister(userRegisterRequest)
	if err != nil {
		res.BadRequest(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	res.Success("REGISTERED", resources.UserResource(user))
}

func ActivateUser(w http.ResponseWriter, r *http.Request) {
	res := rw.ResponseWriter{W: &w}

	activatedKey := mux.Vars(r)["key"]

	user, err := services.ActivateUser(activatedKey)
	if err != nil {
		res.BadRequest(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	res.Success("USER_ACTIVATED", resources.UserResource(user))
}

func Login(w http.ResponseWriter, r *http.Request) {
	res := rw.ResponseWriter{W: &w}

	UserLoginRequest, err := v.UserLoginValidation(r)
	if err != nil {
		res.BadRequest(err.Error())
		return
	}

	user, err := services.Login(UserLoginRequest)
	if err != nil {
		res.BadRequest(err.Error())
		return
	}

	tokenString, err := authentication.Sign(user)
	if err != nil {
		res.BadRequest(err.Error())
		return
	}

	res.SuccessWithToken("LOGGED_IN", tokenString, resources.UserResource(user))
}

func GetUserProfiles(w http.ResponseWriter, r *http.Request) {
	res := rw.ResponseWriter{W: &w}

	userId, _ := strconv.Atoi(mux.Vars(r)["user_id"])

	profiles, err := services.GetUserProfiles(userId)
	if err != nil {
		res.BadRequest(err.Error())
		return
	}

	res.Success("USER_PROFILES", resources.UserProfilesResource(profiles))
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	res := rw.ResponseWriter{W: &w}

	vars := mux.Vars(r)

	userId, _ := strconv.Atoi(vars["user_id"])

	profileId, _ := strconv.Atoi(vars["profile_id"])

	profile, err := services.GetUserProfile(userId, profileId)
	if err != nil {
		res.BadRequest(err.Error())
	}

	res.Success("PROFILE_RETRIEVED", resources.ProfileResource(profile))
}

func ActivateProfile(w http.ResponseWriter, r *http.Request) {
	res := rw.ResponseWriter{W: &w}

	profile, err := services.ActivateUserProfile(r)
	if err != nil {
		res.BadRequest(err.Error())
		return
	}

	res.Success("PROFILE_ACTIVATED", resources.UserProfileResource(profile))
}
