package requests

import "net/http"

type EducationCreateRequest struct {
	Title       string `json:"title" validate:"required"`
	Reference   string `json:"reference" validate:"required"`
	Description string `json:"description" validate:"required"`
	Link        string `json:"link" validate:"required"`
	Date        string `json:"date" validate:"required"`
}

type EducationUpdateRequest struct {
	Title       string `json:"title" validate:"required"`
	Reference   string `json:"reference" validate:"required"`
	Description string `json:"description" validate:"required"`
	Link        string `json:"link" validate:"required"`
	Date        string `json:"date" validate:"required"`
}

func (s *EducationCreateRequest) ValidateRequest(r *http.Request) error {
	return validateRequest(r, s)
}

func (s *EducationUpdateRequest) ValidateRequest(r *http.Request) error {
	return validateRequest(r, s)
}
