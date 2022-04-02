package responses

import (
	"encoding/json"
	"net/http"
)

type response interface {
	send(w *http.ResponseWriter)
}

func ErrorBadRequestResponse(w *http.ResponseWriter, err error) {
	JsonResponse(w, BadRequestResponse{Message: err.Error()})
}

func JsonResponse(w *http.ResponseWriter, res response) {
	res.send(w)
}

type responseData struct {
	Code        int         `json:"code"`
	Status      string      `json:"status"`
	AccessToken string      `json:"access_token,omitempty"`
	TokenType   string      `json:"token_type,omitempty"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	Error       bool        `json:"error"`
	ErrorType   string      `json:"error_type"`
}

// Success response
func (r SuccessResponse) send(w *http.ResponseWriter) {

	res := responseData{
		Code:    http.StatusOK,
		Status:  "success",
		Message: r.Message,
		Data:    r.Data,
		Error:   false,
	}

	res.sendJson(w)
}

// Bad request
func (r BadRequestResponse) send(w *http.ResponseWriter) {

	res := responseData{
		Code:      http.StatusBadRequest,
		Status:    "error",
		Message:   r.Message,
		Data:      make([]interface{}, 0),
		Error:     true,
		ErrorType: "bad_request",
	}

	res.sendJson(w)
}

// Unauthorized
func (r UnauthorizedResponse) send(w *http.ResponseWriter) {

	res := responseData{
		Code:      http.StatusUnauthorized,
		Status:    "error",
		Message:   r.Message,
		Data:      make([]interface{}, 0),
		Error:     true,
		ErrorType: "unauthorized",
	}

	res.sendJson(w)
}

// Success with token string response
func (r SuccessWithTokenResponse) send(w *http.ResponseWriter) {

	res := responseData{
		Code:        http.StatusOK,
		Status:      "success",
		Message:     r.Message,
		AccessToken: r.Token,
		TokenType:   "bearer",
		Data:        r.Data,
		Error:       false,
	}

	res.sendJson(w)
}

// Method not allowed response
func (r MethodNotAllowedResponse) send(w *http.ResponseWriter) {

	res := responseData{
		Code:      http.StatusMethodNotAllowed,
		Status:    "error",
		Message:   r.Message,
		Data:      make([]interface{}, 0),
		Error:     true,
		ErrorType: "not_allowed",
	}

	res.sendJson(w)
}

// Not found http response
func (r NotFoundResponse) send(w *http.ResponseWriter) {

	res := responseData{
		Code:      http.StatusNotFound,
		Status:    "error",
		Message:   r.Message,
		Data:      make([]interface{}, 0),
		Error:     true,
		ErrorType: "not_allowed",
	}

	res.sendJson(w)
}

func setHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
}

func OptionsResponse(w *http.ResponseWriter) {
	setHeaders(w)
	(*w).WriteHeader(http.StatusOK)
}

func (res *responseData) sendJson(w *http.ResponseWriter) {
	setHeaders(w)
	(*w).WriteHeader(res.Code)
	json.NewEncoder(*w).Encode(res.Data)
}
