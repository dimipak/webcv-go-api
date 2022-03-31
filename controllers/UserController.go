package controllers

import (
	"app/externalServices"
	"app/helpers"
	"app/mappers"
	"app/requests"
	"app/resources"
	res "app/responses"
	"app/services"
	"app/systemService/authentication"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	UserRegisterRequest  requests.UserRegisterRequest
	updateProfileRequest requests.UpdateProfileRequest
)

var (
	profileService services.ProfileService
)

func Register(w http.ResponseWriter, r *http.Request) {

	err := UserRegisterRequest.ValidateRequest(r)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	user, err := services.UserRegister(UserRegisterRequest)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "REGISTERED",
		Data:    resources.UserResource(user),
	})
}

func ActivateUser(w http.ResponseWriter, r *http.Request) {

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

func Login(w http.ResponseWriter, r *http.Request) {

	var reqBody requests.UserLoginRequest

	err := reqBody.ValidateRequest(r)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	user, err := services.Login(reqBody)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	tokenString, err := authentication.Sign(user)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessWithTokenResponse{
		Message: "LOGGED_IN",
		Token:   tokenString,
		Data:    resources.UserResource(user),
	})
}

func GetUserProfiles(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(mux.Vars(r)["user_id"])

	profiles, err := services.GetUserProfiles(userId)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
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
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PROFILE_RETRIEVED",
		Data:    resources.ProfileResource(profile),
	})
}

func ActivateProfile(w http.ResponseWriter, r *http.Request) {

	profile, err := services.ActivateUserProfile(r)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PROFILE_ACTIVATED",
		Data:    resources.UserProfileResource(profile),
	})
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(mux.Vars(r)["user_id"])

	var reqBody requests.CreateProfileRequest

	err := reqBody.ValidateRequest(r)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PROFILE_CREATED",
		Data:    resources.ProfileResource(services.CreateProfile(userId, reqBody)),
	})
}

func UploadProfileImage(w http.ResponseWriter, r *http.Request) {

	// var requestBody requests.UploadProfileImageRequest

	// err := requestBody.ValidateRequest(r)
	// if err != nil {
	// 	res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
	// 	return
	// }

	routeIds, err := helpers.RequestRoute(r, "user_id", "profile_id")
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	var profileService services.ProfileService
	profileService.ProfileId = routeIds["profile_id"]

	profile, err := profileService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	if profile.UserId != routeIds["user_id"] {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "wrong profile"})
		return
	}

	img := externalServices.ImageUpload{
		Allowed:  "profile_image",
		Path:     "profile/",
		FileType: "jpg",
	}

	url := img.Upload(r)

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PROFILE_IMAGE_UPDATED",
		Data:    resources.ProfileResource(services.UpdateProfileImage(routeIds["profile_id"], url)),
	})
}

func UploadCoverImage(w http.ResponseWriter, r *http.Request) {

	// var requestBody requests.UploadProfileImageRequest

	// err := requestBody.ValidateRequest(r)
	// if err != nil {
	// 	res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
	// 	return
	// }

	routeIds, err := helpers.RequestRoute(r, "user_id", "profile_id")
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	var profileService services.ProfileService
	profileService.ProfileId = routeIds["profile_id"]

	profile, err := profileService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	if profile.UserId != routeIds["user_id"] {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "wrong profile"})
		return
	}

	img := externalServices.ImageUpload{
		Allowed:  "cover_image",
		Path:     "cover/",
		FileType: "jpg",
	}

	url := img.Upload(r)

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PROFILE_IMAGE_UPDATED",
		Data:    resources.ProfileResource(services.UpdateProfileImage(routeIds["profile_id"], url)),
	})
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	// Validate request
	err := updateProfileRequest.ValidateRequest(r)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	route := helpers.RouteParams(r)

	fmt.Println("globals. user_id", route.UserId)
	fmt.Println("globals. profile_id", route.ProfileId)
	// Ger url ids
	routeIds, err := helpers.RequestRoute(r, "user_id", "profile_id", "Asd")
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	// Get profile
	_, err = profileService.SetProfileId(routeIds["profile_id"]).SetUserId(routeIds["user_id"]).GetUserProfile()
	// Check if profile is for correct user
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	// Update profile
	updatedProfile, err := profileService.SetProfileUpdateRequest(updateProfileRequest).UpdateById(mappers.UpdateProfileMapper)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	// success response
	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PROFILE_UPDATED",
		Data:    resources.ProfileResource(updatedProfile.SocialNetworks()),
	})
}

func GetUserProfileSkills(w http.ResponseWriter, r *http.Request) {

	routeIds, err := helpers.RequestRoute(r, "user_id", "profile_id")
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot parse url"})
		return
	}

	var profileService services.ProfileService
	var skillService services.SkillService

	profileService.ProfileId = routeIds["profile_id"]
	skillService.ProfileId = routeIds["profile_id"]

	profile, err := profileService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot parse url"})
		return
	}

	if profile.UserId != routeIds["user_id"] {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Wrong profile"})
		return
	}

	skills, err := skillService.GetByProfileId()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "SKILLS_RETRIEVED!!",
		Data:    resources.SkillsResources(skills),
	})
}

func GetUserProfileSkill(w http.ResponseWriter, r *http.Request) {

	routeIds, err := helpers.RequestRoute(r, "user_id", "profile_id", "skill_id")
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot parse url"})
		return
	}

	var profileService services.ProfileService
	var skillService services.SkillService

	profileService.ProfileId = routeIds["profile_id"]
	profile, err := profileService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot get profile"})
		return
	}

	if profile.UserId != routeIds["user_id"] {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Wrong profile"})
		return
	}

	skillService.SkillId = routeIds["skill_id"]
	skill, err := skillService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "SKILL_RETRIEVED",
		Data:    resources.SkillResources(skill),
	})
}

func UpdateUserProfileSkill(w http.ResponseWriter, r *http.Request) {

	var requestBody requests.SkillUpdateRequest

	err := requests.ValidateRequest(r, &requestBody)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	routeIds, err := helpers.RequestRoute(r, "user_id", "profile_id", "skill_id")
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot parse url"})
		return
	}

	var profileService services.ProfileService
	var skillService services.SkillService

	profileService.ProfileId = routeIds["profile_id"]
	profile, err := profileService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot get profile"})
		return
	}

	if profile.UserId != routeIds["user_id"] {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Wrong profile"})
		return
	}

	skillService.SkillId = routeIds["skill_id"]
	skillService.UpdateSkillRequestBody = requestBody
	skill, err := skillService.UpdateById(mappers.UpdateSkillMapper)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "SKILL_UPDATED",
		Data:    resources.SkillResources(skill),
	})
}

func CreateUserProfileSkill(w http.ResponseWriter, r *http.Request) {

	var requestBody requests.SkillCreateRequest

	err := requestBody.ValidateRequest(r)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	routeIds, err := helpers.RequestRoute(r, "user_id", "profile_id")
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot parse url"})
		return
	}

	var profileService services.ProfileService
	var skillService services.SkillService

	profileService.ProfileId = routeIds["profile_id"]
	profile, err := profileService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot get profile"})
		return
	}

	if profile.UserId != routeIds["user_id"] {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Wrong profile"})
		return
	}

	skillService.CreateSkillRequestBody = requestBody
	skillService.ProfileId = routeIds["profile_id"]
	skill, err := skillService.Create(mappers.CreateSkillMapper)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "SKILL_CREATED",
		Data:    resources.SkillResources(skill),
	})
}

func DeleteUserProfileSkill(w http.ResponseWriter, r *http.Request) {

	routeIds, err := helpers.RequestRoute(r, "user_id", "profile_id", "skill_id")
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot parse url"})
		return
	}

	var profileService services.ProfileService
	var skillService services.SkillService

	profileService.ProfileId = routeIds["profile_id"]
	profile, err := profileService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot get profile"})
		return
	}

	if profile.UserId != routeIds["user_id"] {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Wrong profile"})
		return
	}

	skillService.SkillId = routeIds["skill_id"]
	skill, err := skillService.DeleteById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "SKILL_DELETED",
		Data:    resources.SkillResources(skill),
	})
}

func UpdateUserProfileSkillsOrder(w http.ResponseWriter, r *http.Request) {

	var requestBody requests.SkillUpdateOrderRequest

	err := requestBody.ValidateRequest(r)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	routeIds, err := helpers.RequestRoute(r, "user_id", "profile_id")
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot parse url"})
		return
	}

	var profileService services.ProfileService
	var skillService services.SkillService

	profileService.ProfileId = routeIds["profile_id"]
	profile, err := profileService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot get profile"})
		return
	}

	if profile.UserId != routeIds["user_id"] {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Wrong profile"})
		return
	}

	skillService.ProfileId = routeIds["profile_id"]
	skills, err := skillService.UpdateSkillsOrder(requestBody.SkillIds)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: err.Error()})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "SKILLS_ORDER_CHANGED",
		Data:    resources.SkillsResources(skills),
	})
}
