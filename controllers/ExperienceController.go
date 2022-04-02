package controllers

import (
	"app/mappers"
	"app/models"
	"app/resources"
	res "app/responses"
	"app/system"
	"net/http"
)

type ExperienceController struct{}

func (e *ExperienceController) Get(w http.ResponseWriter, r *http.Request) {

	routes := system.RouteParams(r)

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EXPERIENCES_RETRIEVED",
		Data:    resources.ExperiencesResources(profile.Experiences().Experience),
	})
}

func (e *ExperienceController) Show(w http.ResponseWriter, r *http.Request) {

	routes := system.RouteParams(r)

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	experience, err := profile.Experiences().Experience.GetExperience(routes.ExperienceId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EXPERIENCE_RETRIEVED",
		Data:    resources.ExperienceResources(experience),
	})
}

func (e *ExperienceController) Create(w http.ResponseWriter, r *http.Request) {
	err := experienceCreateRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	routes := system.RouteParams(r)

	_, err = profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	experience, err := experienceService.SetProfileId(routes.ProfileId).SetCreateRequest(experienceCreateRequest).Create(mappers.ExperienceCreateMapper)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EXPERIENCE_CREATED",
		Data:    resources.ExperienceResources(experience),
	})
}

func (e *ExperienceController) Update(w http.ResponseWriter, r *http.Request) {
	err := experienceUpdateRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	routes := system.RouteParams(r)

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	experience, err := profile.Experiences().Experience.GetExperience(routes.ExperienceId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	experience, err = experienceService.SetProfileId(routes.ProfileId).SetExperience(experience).Update(models.Experience{
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

func (e *ExperienceController) Delete(w http.ResponseWriter, r *http.Request) {

	routes := system.RouteParams(r)

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	experience, err := profile.Experiences().Experience.GetExperience(routes.ExperienceId)
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
