package requests

import (
	"net/http"
)

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s *UserRegisterRequest) ValidateRequest(r *http.Request) error {
	return validateRequest(r, s)
}
