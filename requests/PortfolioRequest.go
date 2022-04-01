package requests

import "net/http"

type PortfolioCreateRequest struct {
	Name       string `json:"name" validate:"required"`
	Type       string `json:"type" validate:"required"`
	Technology string `json:"technology" validate:"required"`
	Customer   string `json:"customer" validate:"required"`
	WebsiteUrl string `json:"website_url" validate:"required"`
}

type PortfolioUpdateRequest struct {
	Name       string `json:"name" validate:"required"`
	Type       string `json:"type" validate:"required"`
	Technology string `json:"technology" validate:"required"`
	Customer   string `json:"customer" validate:"required"`
	WebsiteUrl string `json:"website_url" validate:"required"`
}

func (s *PortfolioCreateRequest) ValidateRequest(r *http.Request) error {
	return validateRequest(r, s)
}

func (s *PortfolioUpdateRequest) ValidateRequest(r *http.Request) error {
	return validateRequest(r, s)
}
