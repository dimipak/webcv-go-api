package validations

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func UserRegisterValidation(r *http.Request) (UserRegisterRequest, error) {

	var req UserRegisterRequest

	json.NewDecoder(r.Body).Decode(&req)

	request := UserRegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	validate = validator.New()

	return req, validate.Struct(request)
}
