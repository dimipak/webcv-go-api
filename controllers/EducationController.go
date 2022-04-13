package controllers

import (
	"app/mappers"
	"app/requests"
	"app/resources"
	res "app/responses"
	"app/services"
	"app/system"
	"net/http"
)

type EducationController struct{}

func (e *EducationController) Get(w http.ResponseWriter, r *http.Request) {

	routes := system.RouteParams(r)

	var profileService services.ProfileService

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EDUCATIONS_RETRIEVED",
		Data:    resources.EducationsResources(profile.GetEducations().Education),
	})
}

func (e *EducationController) Show(w http.ResponseWriter, r *http.Request) {

	routes := system.RouteParams(r)

	var profileService services.ProfileService

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	education, err := profile.GetEducations().Education.GetOne(routes.EducationId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "EDUCATION_RETRIEVED",
		Data:    resources.EducationResources(education),
	})
}

func (e *EducationController) Create(w http.ResponseWriter, r *http.Request) {
	var educationCreateRequest requests.EducationCreateRequest

	err := educationCreateRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	routes := system.RouteParams(r)

	var profileService services.ProfileService
	var educationService services.EducationService

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

func (e *EducationController) Update(w http.ResponseWriter, r *http.Request) {
	var educationUpdateRequest requests.EducationUpdateRequest

	err := educationUpdateRequest.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	routes := system.RouteParams(r)
	var profileService services.ProfileService

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	education, err := profile.GetEducations().Education.GetOne(routes.EducationId)
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

func (e *EducationController) Delete(w http.ResponseWriter, r *http.Request) {
	routes := system.RouteParams(r)
	var profileService services.ProfileService

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	education, err := profile.GetEducations().Education.GetOne(routes.EducationId)
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
