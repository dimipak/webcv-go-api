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

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	res.JsonResponse(&w, res.SuccessResponse{
		Message: "SKILLS_RETRIEVED",
		Data:    resources.SkillsResources(profile.GetSkills().Skills.OrderByOrder()),
	})
}

func (s *SkillController) Show(w http.ResponseWriter, r *http.Request) {

	routes := system.RouteParams(r)

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	skill, err := profile.GetSkills().Skills.GetById(routes.SkillId)
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

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	skill, err := profile.GetSkills().Skills.GetById(routes.SkillId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	skill.Update(mappers.UpdateSkillMapper(requestBody))

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

	var profileService services.ProfileService
	var skillService services.SkillService

	routes := system.RouteParams(r)

	_, err = profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
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

	profile, err := profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	skill, err := profile.GetSkills().Skills.GetById(routes.SkillId)
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
		return
	}

	if skill.Delete() != nil {
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

	_, err = profileService.SetProfileId(routes.ProfileId).SetUserId(routes.UserId).GetUserProfile()
	if err != nil {
		res.ErrorBadRequestResponse(&w, err)
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
