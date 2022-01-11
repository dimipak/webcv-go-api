package validations

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func UserLoginValidation(r *http.Request) (UserLoginRequest, error) {

	var req UserLoginRequest

	json.NewDecoder(r.Body).Decode(&req)

	request := UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	validate = validator.New()

	return req, validate.Struct(request)
}
