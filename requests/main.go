package requests

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type validateReq interface {
	ValidateRequest(r *http.Request) error
}

func ValidateRequest(r *http.Request, req validateReq) error {
	return req.ValidateRequest(r)
}

func validateRequest(r *http.Request, rb interface{}) error {

	json.NewDecoder(r.Body).Decode(rb)

	validate = validator.New()

	return validate.Struct(rb)
}
