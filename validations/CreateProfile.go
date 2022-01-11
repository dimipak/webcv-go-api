package validations

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type CreateProfileRequest struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func CreateProfileValidation(r *http.Request) (CreateProfileRequest, error) {
	var req CreateProfileRequest

	json.NewDecoder(r.Body).Decode(&req)

	validate = validator.New()

	return req, validate.Struct(req)
}
