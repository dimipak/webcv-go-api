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

type SkillController struct{}

func (s *SkillController) Get(w http.ResponseWriter, r *http.Request) {

	routes := system.RouteParams(r)

	var profileService services.ProfileService
	var skillService services.SkillService

	profileService.ProfileId = routes.ProfileId
	skillService.ProfileId = routes.ProfileId

	profile, err := profileService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot parse url"})
		return
	}

	if profile.UserId != routes.UserId {
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

func (s *SkillController) Show(w http.ResponseWriter, r *http.Request) {

	routes := system.RouteParams(r)

	var profileService services.ProfileService
	var skillService services.SkillService

	profileService.ProfileId = routes.ProfileId
	profile, err := profileService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot get profile"})
		return
	}

	if profile.UserId != routes.UserId {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Wrong profile"})
		return
	}

	skillService.SkillId = routes.SkillId
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

func (s *SkillController) Update(w http.ResponseWriter, r *http.Request) {

	var requestBody requests.SkillUpdateRequest

	err := requests.ValidateRequest(r, &requestBody)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	routes := system.RouteParams(r)

	var profileService services.ProfileService
	var skillService services.SkillService

	profileService.ProfileId = routes.ProfileId
	profile, err := profileService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot get profile"})
		return
	}

	if profile.UserId != routes.UserId {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Wrong profile"})
		return
	}

	skillService.SkillId = routes.SkillId
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

func (s *SkillController) Create(w http.ResponseWriter, r *http.Request) {

	var requestBody requests.SkillCreateRequest

	err := requestBody.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	routes := system.RouteParams(r)

	var profileService services.ProfileService
	var skillService services.SkillService

	profileService.ProfileId = routes.ProfileId
	profile, err := profileService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot get profile"})
		return
	}

	if profile.UserId != routes.UserId {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Wrong profile"})
		return
	}

	skillService.CreateSkillRequestBody = requestBody
	skillService.ProfileId = routes.ProfileId
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

func (s *SkillController) Delete(w http.ResponseWriter, r *http.Request) {

	routes := system.RouteParams(r)

	var profileService services.ProfileService
	var skillService services.SkillService

	profileService.ProfileId = routes.ProfileId
	profile, err := profileService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot get profile"})
		return
	}

	if profile.UserId != routes.UserId {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Wrong profile"})
		return
	}

	skillService.SkillId = routes.SkillId
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

func (s *SkillController) UpdateOrder(w http.ResponseWriter, r *http.Request) {

	var requestBody requests.SkillUpdateOrderRequest

	err := requestBody.ValidateRequest(r)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	routes := system.RouteParams(r)

	var profileService services.ProfileService
	var skillService services.SkillService

	profileService.ProfileId = routes.ProfileId
	profile, err := profileService.GetById()
	if err != nil {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Cannot get profile"})
		return
	}

	if profile.UserId != routes.UserId {
		res.JsonResponse(&w, res.BadRequestResponse{Message: "Wrong profile"})
		return
	}

	skillService.ProfileId = routes.ProfileId
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
