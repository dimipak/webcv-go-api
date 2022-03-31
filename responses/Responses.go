package responses

import "net/http"

type SuccessResponse struct {
	Message string
	Data    interface{}
}

type BadRequestResponse struct {
	Message string
}

type UnauthorizedResponse struct {
	W *http.ResponseWriter
}

type SuccessWithTokenResponse struct {
	Message string
	Token   string
	Data    interface{}
}

type MethodNotAllowedResponse struct {
	Message string
}

type NotFoundResponse struct {
	Message string
}

type OptionsResponse struct {
	W *http.ResponseWriter
}
