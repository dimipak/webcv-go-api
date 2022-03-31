package requests

import (
	"net/http"
)

type CreateProfileRequest struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func (u *CreateProfileRequest) ValidateRequest(r *http.Request) error {
	return validateRequest(r, u)
}
