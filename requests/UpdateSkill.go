package requests

import "net/http"

type SkillUpdateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Progress    int    `json:"progress" validate:"required"`
}

type SkillCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Progress    int    `json:"progress" validate:"required"`
}

type SkillUpdateOrderRequest struct {
	SkillIds []int `json:"skill_ids" validate:"required"`
}

func (s *SkillUpdateRequest) ValidateRequest(r *http.Request) error {
	return validateRequest(r, s)
}

func (s *SkillCreateRequest) ValidateRequest(r *http.Request) error {
	return validateRequest(r, s)
}

func (s *SkillUpdateOrderRequest) ValidateRequest(r *http.Request) error {
	return validateRequest(r, s)
}
