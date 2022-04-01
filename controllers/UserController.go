package controllers

import (
	"app/externalServices"
	"app/helpers"
	"app/mappers"
	"app/models"
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

// Request
var (
	UserRegisterRequest     requests.UserRegisterRequest
	updateProfileRequest    requests.UpdateProfileRequest
	portfolioCreateRequest  requests.PortfolioCreateRequest
	portfolioUpdateRequest  requests.PortfolioUpdateRequest
	experienceCreateRequest requests.ExperienceCreateRequest
	experienceUpdateRequest requests.ExperienceUpdateRequest
	educationCreateRequest  requests.EducationCreateRequest
	educationUpdateRequest  requests.EducationUpdateRequest
)

// Services
var (
	profileService    services.ProfileService
	portfolioService  services.PortfolioService
	experienceService services.ExperienceService
	educationService  services.EducationService
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
		res.ErrorBadRequestResponse(&w, err)
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
		res.ErrorBadRequestResponse(&w, err)
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
		res.ErrorBadRequestResponse(&w, err)
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
		res.ErrorBadRequestResponse(&w, err)
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
	// 	res.ErrorBadRequestResponse(&w, err)
	// 	return
	// }

	routeIds, err := helpers.RequestRoute(r, "user_id", "profile_id")
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	var profileService services.ProfileService
	profileService.ProfileId = routeIds["profile_id"]

	profile, err := profileService.GetById()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
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

	url, err := img.Upload(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PROFILE_IMAGE_UPDATED",
		Data:    resources.ProfileResource(services.UpdateProfileImage(routeIds["profile_id"], url)),
	})
}

func UploadCoverImage(w http.ResponseWriter, r *http.Request) {

	// var requestBody requests.UploadProfileImageRequest

	// err := requestBody.ValidateRequest(r)
	// if err != nil {
	// 	res.ErrorBadRequestResponse(&w, err)
	// 	return
	// }

	profile, err := profileService.SetUserId(helpers.RouteParams(r).UserId).SetProfileId(helpers.RouteParams(r).ProfileId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	img := externalServices.ImageUpload{
		Allowed:  "cover_image",
		Path:     "cover/",
		FileType: "jpg",
	}

	url, err := img.Upload(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	err = profile.Update(models.Profile{CoverImage: url})
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PROFILE_IMAGE_UPDATED",
		Data:    resources.ProfileResource(profile),
	})
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	// Validate request
	err := updateProfileRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	// Get profile
	_, err = profileService.SetProfileId(helpers.RouteParams(r).ProfileId).SetUserId(helpers.RouteParams(r).UserId).GetUserProfile()
	// Check if profile is for correct user
	if err != nil {
		fmt.Println("here")
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	// Update profile
	updatedProfile, err := profileService.SetProfileUpdateRequest(updateProfileRequest).UpdateById(mappers.UpdateProfileMapper)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
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
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "SKILLS_RETRIEVED",
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
		res.ErrorBadRequestResponse(&w, err)
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
		res.ErrorBadRequestResponse(&w, err)
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
		res.ErrorBadRequestResponse(&w, err)
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
		res.ErrorBadRequestResponse(&w, err)
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
		res.ErrorBadRequestResponse(&w, err)
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
		res.ErrorBadRequestResponse(&w, err)
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
		res.ErrorBadRequestResponse(&w, err)
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
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "SKILLS_ORDER_CHANGED",
		Data:    resources.SkillsResources(skills),
	})
}

func GetUserProfilePortfolios(w http.ResponseWriter, r *http.Request) {

	profile, err := profileService.SetProfileId(helpers.RouteParams(r).ProfileId).SetUserId(helpers.RouteParams(r).UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PORTFOLIO_RETRIEVED",
		Data:    resources.PortfoliosResources(profile.Portfolios().Portfolio),
	})
}

func GetUserProfilePortfolio(w http.ResponseWriter, r *http.Request) {

	profile, err := profileService.SetProfileId(helpers.RouteParams(r).ProfileId).SetUserId(helpers.RouteParams(r).UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio, err := profile.Portfolios().Portfolio.GetOne(helpers.RouteParams(r).PortfolioId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PORTFOLIO_RETRIEVED",
		Data:    resources.PortfolioResources(portfolio),
	})
}

func CreateUserProfilePortfolio(w http.ResponseWriter, r *http.Request) {

	err := portfolioCreateRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	_, err = profileService.SetProfileId(helpers.RouteParams(r).ProfileId).SetUserId(helpers.RouteParams(r).UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio, err := portfolioService.SetProfileId(helpers.RouteParams(r).ProfileId).SetPortfolioCreateRequest(portfolioCreateRequest).Create(mappers.PortfolioCreateMapper)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PORTFOLIO_CREATED",
		Data:    resources.PortfolioResources(portfolio),
	})
}

func UpdateUserPorfilePortfolio(w http.ResponseWriter, r *http.Request) {

	err := portfolioUpdateRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	profile, err := profileService.SetProfileId(helpers.RouteParams(r).ProfileId).SetUserId(helpers.RouteParams(r).UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio, err := profile.Portfolios().Portfolio.GetOne(helpers.RouteParams(r).PortfolioId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio.Update(models.Portfolio{
		Name:       portfolioUpdateRequest.Name,
		Type:       portfolioUpdateRequest.Type,
		Technology: portfolioUpdateRequest.Technology,
		Customer:   portfolioUpdateRequest.Customer,
		WebsiteUrl: portfolioUpdateRequest.WebsiteUrl,
	})

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PORTFOLIO_UPDATED",
		Data:    resources.PortfolioResources(portfolio),
	})
}

func UploadPortfolioImage(w http.ResponseWriter, r *http.Request) {

	profile, err := profileService.SetProfileId(helpers.RouteParams(r).ProfileId).SetUserId(helpers.RouteParams(r).UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio, err := profile.Portfolios().Portfolio.GetOne(helpers.RouteParams(r).PortfolioId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	img := externalServices.ImageUpload{
		Allowed:  "portfolio_image",
		Path:     "portfolio/",
		FileType: "jpg",
	}

	url, err := img.Upload(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio.Update(models.Portfolio{
		ImageUrl: url,
	})

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PORTFOLIO_IMAGE_UPDATEDD",
		Data:    resources.PortfolioResources(portfolio),
	})
}

func DeleteUserProfilePortfolio(w http.ResponseWriter, r *http.Request) {

	profile, err := profileService.SetProfileId(helpers.RouteParams(r).ProfileId).SetUserId(helpers.RouteParams(r).UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio, err := profile.Portfolios().Portfolio.GetOne(helpers.RouteParams(r).PortfolioId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	portfolio.Delete()

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PORTFOLIO_DELETED",
		Data:    resources.PortfolioResources(portfolio),
	})
}

func GetUserProfileExperiences(w http.ResponseWriter, r *http.Request) {

	profile, err := profileService.SetProfileId(helpers.RouteParams(r).ProfileId).SetUserId(helpers.RouteParams(r).UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EXPERIENCES_RETRIEVED",
		Data:    resources.ExperiencesResources(profile.Experiences().Experience),
	})
}

func GetUserProfileExperience(w http.ResponseWriter, r *http.Request) {

	profile, err := profileService.SetProfileId(helpers.RouteParams(r).ProfileId).SetUserId(helpers.RouteParams(r).UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	experience, err := profile.Experiences().Experience.GetExperience(helpers.RouteParams(r).ExperienceId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EXPERIENCE_RETRIEVED",
		Data:    resources.ExperienceResources(experience),
	})
}

func CreateUserProfileExperience(w http.ResponseWriter, r *http.Request) {
	err := experienceCreateRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	_, err = profileService.SetProfileId(helpers.RouteParams(r).ProfileId).SetUserId(helpers.RouteParams(r).UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	experience, err := experienceService.SetProfileId(helpers.RouteParams(r).ProfileId).SetCreateRequest(experienceCreateRequest).Create(mappers.ExperienceCreateMapper)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EXPERIENCE_CREATED",
		Data:    resources.ExperienceResources(experience),
	})
}

func UpdateUserProfileExperience(w http.ResponseWriter, r *http.Request) {
	err := experienceUpdateRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	profile, err := profileService.SetProfileId(helpers.RouteParams(r).ProfileId).SetUserId(helpers.RouteParams(r).UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	experience, err := profile.Experiences().Experience.GetExperience(helpers.RouteParams(r).ExperienceId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	experience, err = experienceService.SetProfileId(helpers.RouteParams(r).ProfileId).SetExperience(experience).Update(models.Experience{
		CompanyName: experienceUpdateRequest.CompanyName,
		Role:        experienceUpdateRequest.Role,
		Description: experienceUpdateRequest.Description,
		Country:     experienceUpdateRequest.Country,
		City:        experienceUpdateRequest.City,
		StartDate:   experienceUpdateRequest.StartDate,
		EndDate:     experienceUpdateRequest.EndDate,
	})
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EXPERIENCE_UPDATED",
		Data:    resources.ExperienceResources(experience),
	})
}

func DeleteUserProfileExperience(w http.ResponseWriter, r *http.Request) {
	profile, err := profileService.SetProfileId(helpers.RouteParams(r).ProfileId).SetUserId(helpers.RouteParams(r).UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	experience, err := profile.Experiences().Experience.GetExperience(helpers.RouteParams(r).ExperienceId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	err = experience.Delete()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EXPERIENCE_DELETED",
		Data:    resources.ExperienceResources(experience),
	})
}

func GetUserProfileEducations(w http.ResponseWriter, r *http.Request) {

	routes := helpers.RouteParams(r)

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	educations := profile.Educations().Education

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EDUCATIONS_RETRIEVED",
		Data:    resources.EducationsResources(educations),
	})
}

func GetUserProfileEducation(w http.ResponseWriter, r *http.Request) {

	routes := helpers.RouteParams(r)

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	education, err := profile.Educations().Education.GetOne(routes.EducationId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EDUCATION_RETRIEVED",
		Data:    resources.EducationResources(education),
	})
}

func CreateUserProfileEducation(w http.ResponseWriter, r *http.Request) {

	err := educationCreateRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	routes := helpers.RouteParams(r)

	_, err = profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	education, err := educationService.SetProfileId(routes.ProfileId).SetCreateRequest(educationCreateRequest).Create(mappers.EducationCreateMapper)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EDUCATION_CREATED",
		Data:    resources.EducationResources(education),
	})
}

func UpdateUserProfileEducation(w http.ResponseWriter, r *http.Request) {

	err := educationUpdateRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	routes := helpers.RouteParams(r)

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	education, err := profile.Educations().Education.GetOne(routes.EducationId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	education.Update(mappers.EducationUpdateMapper(educationUpdateRequest))

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EDUCATION_UPDATED",
		Data:    resources.EducationResources(education),
	})
}

func DeleteUserProfileEducation(w http.ResponseWriter, r *http.Request) {
	routes := helpers.RouteParams(r)

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	education, err := profile.Educations().Education.GetOne(routes.EducationId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	education.Delete()

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EDUCATION_DELETED",
		Data:    resources.EducationResources(education),
	})
}
