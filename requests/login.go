package requests

import (
	"net/http"
)

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (u *UserLoginRequest) ValidateRequest(r *http.Request) error {
	return validateRequest(r, u)
}
