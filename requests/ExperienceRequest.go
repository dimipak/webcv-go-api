package requests

import "net/http"

type ExperienceCreateRequest struct {
	CompanyName string `json:"company_name" validate:"required"`
	Role        string `json:"role" validate:"required"`
	Description string `json:"description" validate:"required"`
	Country     string `json:"country" validate:"required"`
	City        string `json:"city" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
}

type ExperienceUpdateRequest struct {
	CompanyName string `json:"company_name" validate:"required"`
	Role        string `json:"role" validate:"required"`
	Description string `json:"description" validate:"required"`
	Country     string `json:"country" validate:"required"`
	City        string `json:"city" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
}

func (s *ExperienceCreateRequest) ValidateRequest(r *http.Request) error {
	return validateRequest(r, s)
}

func (s *ExperienceUpdateRequest) ValidateRequest(r *http.Request) error {
	return validateRequest(r, s)
}