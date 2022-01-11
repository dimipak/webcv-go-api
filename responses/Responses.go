package responses

import "net/http"

type SuccessResponse struct {
	W       *http.ResponseWriter
	Message string
	Data    interface{}
}

type BadRequestResponse struct {
	W       *http.ResponseWriter
	Message string
}

type UnauthorizedResponse struct {
	W *http.ResponseWriter
}

type SuccessWithTokenResponse struct {
	W       *http.ResponseWriter
	Message string
	Token   string
	Data    interface{}
}

type MethodNotAllowedResponse struct {
	W       *http.ResponseWriter
	Message string
}

type NotFoundResponse struct {
	W       *http.ResponseWriter
	Message string
}
