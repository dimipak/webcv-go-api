package responses

import (
	"encoding/json"
	"net/http"
)

type response interface {
	send()
}

func JsonResponse(res response) {
	res.send()
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
func (r SuccessResponse) send() {

	res := responseData{
		Code:    http.StatusOK,
		Status:  "success",
		Message: r.Message,
		Data:    r.Data,
		Error:   false,
	}

	sendJson(http.StatusOK, r.W, res)
}

// Bad request
func (r BadRequestResponse) send() {

	res := responseData{
		Code:      http.StatusBadRequest,
		Status:    "error",
		Message:   r.Message,
		Data:      make([]interface{}, 0),
		Error:     true,
		ErrorType: "bad_request",
	}

	sendJson(http.StatusBadRequest, r.W, res)
}

// Unauthorized
func (r UnauthorizedResponse) send() {

	res := responseData{
		Code:      http.StatusUnauthorized,
		Status:    "error",
		Message:   "User is not authorized",
		Data:      make([]interface{}, 0),
		Error:     true,
		ErrorType: "unauthorized",
	}

	sendJson(http.StatusUnauthorized, r.W, res)
}

// Success with token string response
func (r SuccessWithTokenResponse) send() {

	res := responseData{
		Code:        http.StatusOK,
		Status:      "success",
		Message:     r.Message,
		AccessToken: r.Token,
		TokenType:   "bearer",
		Data:        r.Data,
		Error:       false,
	}

	sendJson(http.StatusOK, r.W, res)
}

// Method not allowed response
func (r MethodNotAllowedResponse) send() {

	res := responseData{
		Code:      http.StatusMethodNotAllowed,
		Status:    "error",
		Message:   r.Message,
		Data:      make([]interface{}, 0),
		Error:     true,
		ErrorType: "not_allowed",
	}

	sendJson(http.StatusMethodNotAllowed, r.W, res)
}

// Not found http response
func (r NotFoundResponse) send() {

	res := responseData{
		Code:      http.StatusNotFound,
		Status:    "error",
		Message:   r.Message,
		Data:      make([]interface{}, 0),
		Error:     true,
		ErrorType: "not_allowed",
	}

	sendJson(http.StatusNotFound, r.W, res)
}

func sendJson(status int, w *http.ResponseWriter, data responseData) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(status)
	json.NewEncoder(*w).Encode(data)
}
