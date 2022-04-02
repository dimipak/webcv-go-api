package controllers

import (
	"app/mappers"
	"app/models"
	u "app/pdfGenerator"
	"app/requests"
	"app/resources"
	res "app/responses"
	"app/services"
	"app/system"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type ProfileController struct{}

func (p *ProfileController) GetActiveProfileInfo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	profileId, _ := strconv.Atoi(vars["profile_id"])

	profile, err := services.GetActiveProfileInfo(profileId)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "PROFILE_RETRIEVED",
		Data:    resources.ProfileResource(profile),
	})
}

func (p *ProfileController) GetActiveProfile(w http.ResponseWriter, r *http.Request) {

	profile, err := services.GetActiveProfile()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "ACTIVE_PROFILE_RETRIEVED",
		Data:    resources.ActiveProfileResource(profile),
	})
}

func (p *ProfileController) GetActiveProfileSkills(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	profileId, _ := strconv.Atoi(vars["profile_id"])

	skills, err := services.GetActiveProfileSkills(profileId)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "ACTIVE_PROFILE_SKILLS",
		Data:    resources.SkillsResources(skills),
	})
}

func (p *ProfileController) GetActiveProfilePortfolio(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	profileId, _ := strconv.Atoi(vars["profile_id"])

	portfolio, err := services.GetActiveProfilePortfolio(profileId)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "ACTIVE_PROFILE_PORTFOLIO",
		Data:    resources.PortfoliosResources(portfolio),
	})
}

func (p *ProfileController) GetActiveProfileExperiences(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	profileId, _ := strconv.Atoi(vars["profile_id"])

	experiences, err := services.GetActiveProfileExperiences(profileId)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "ACTIVE_PROFILE_EXPERIENCES",
		Data:    resources.ExperiencesResources(experiences),
	})
}

func (p *ProfileController) GetActiveProfileEducations(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	profileId, _ := strconv.Atoi(vars["profile_id"])

	educations, err := services.GetActiveProfileEducations(profileId)
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{
			Message: err.Error(),
		})
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "ACTIVE_PROFILE_EDUCATIONS",
		Data:    resources.EducationsResources(educations),
	})
}

func (p *ProfileController) CreatePdf(w http.ResponseWriter, r *http.Request) {

	pdf := u.NewRequestPdf("")

	//html template path
	templatePath := "templates/template.html"

	//path for download pdf
	outputPath := "storage/example.pdf"

	//html template data
	templateData := struct {
		Title       string
		Description string
		Company     string
		Contact     string
		Country     string
	}{
		Title:       "HTML to PDF generator",
		Description: "This is the simple HTML to PDF file.",
		Company:     "Jhon Lewis",
		Contact:     "Maria Anders",
		Country:     "Germany",
	}

	if err := pdf.ParseTemplate(templatePath, templateData); err == nil {
		ok, _ := pdf.GeneratePDF(outputPath)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println("ERROR: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(templateData)
}

func (p *ProfileController) Get(w http.ResponseWriter, r *http.Request) {

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

func (p *ProfileController) Show(w http.ResponseWriter, r *http.Request) {

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

func (p *ProfileController) Activate(w http.ResponseWriter, r *http.Request) {

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

func (p *ProfileController) Create(w http.ResponseWriter, r *http.Request) {

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

func (p *ProfileController) UploadProfileImage(w http.ResponseWriter, r *http.Request) {

	// var requestBody requests.UploadProfileImageRequest

	// err := requestBody.ValidateRequest(r)
	// if err != nil {
	// 	res.ErrorBadRequestResponse(&w, err)
	// 	return
	// }
	routes := system.RouteParams(r)

	var profileService services.ProfileService
	profileService.ProfileId = routes.ProfileId

	profile, err := profileService.GetById()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	if profile.UserId != routes.UserId {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "wrong profile"})
		return
	}

	img := system.ImageUpload{
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
		Data:    resources.ProfileResource(services.UpdateProfileImage(routes.ProfileId, url)),
	})
}

func (p *ProfileController) UploadCoverImage(w http.ResponseWriter, r *http.Request) {

	// var requestBody requests.UploadProfileImageRequest

	// err := requestBody.ValidateRequest(r)
	// if err != nil {
	// 	res.ErrorBadRequestResponse(&w, err)
	// 	return
	// }
	routes := system.RouteParams(r)

	profile, err := profileService.SetUserId(routes.UserId).SetProfileId(routes.ProfileId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	img := system.ImageUpload{
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

func (p *ProfileController) Update(w http.ResponseWriter, r *http.Request) {

	// Validate request
	err := updateProfileRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	routes := system.RouteParams(r)

	// Get profile
	_, err = profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
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
