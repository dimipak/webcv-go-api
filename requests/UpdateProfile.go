package requests

import (
	"net/http"
)

type SocialNetwork struct {
	Linkedin string `json:"linkedin" validate:"required"`
	Github   string `json:"github" validate:"required"`
}

type UpdateProfileRequest struct {
	Username      string        `json:"username" validate:"required"`
	Email         string        `json:"email" validate:"required"`
	FirstName     string        `json:"first_name" validate:"required"`
	LastName      string        `json:"last_name" validate:"required"`
	FirstQuote    string        `json:"first_quote" validate:"required"`
	SecondQuote   string        `json:"second_quote" validate:"required"`
	About         string        `json:"about" validate:"required"`
	SocialNetwork SocialNetwork `json:"social_networks" validate:"required"`
}

func (s *UpdateProfileRequest) ValidateRequest(r *http.Request) error {
	return validateRequest(r, s)
}
