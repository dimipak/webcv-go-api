package responses

import (
	"encoding/json"
	"net/http"
)

type response interface {
	send(w *http.ResponseWriter)
}

// func JsonResponse(res response) {
// 	res.send()
// }

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

	sendJson(http.StatusOK, w, res)
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

	sendJson(http.StatusBadRequest, w, res)
}

// Unauthorized
func (r UnauthorizedResponse) send(w *http.ResponseWriter) {

	res := responseData{
		Code:      http.StatusUnauthorized,
		Status:    "error",
		Message:   "User is not authorized",
		Data:      make([]interface{}, 0),
		Error:     true,
		ErrorType: "unauthorized",
	}

	sendJson(http.StatusUnauthorized, w, res)
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

	sendJson(http.StatusOK, w, res)
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

	sendJson(http.StatusMethodNotAllowed, w, res)
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

	sendJson(http.StatusNotFound, w, res)
}

func (r OptionsResponse) send(w *http.ResponseWriter) {

	(*r.W).Header().Set("Content-Type", "application/json")
	(*r.W).Header().Set("Access-Control-Allow-Origin", "*")
	(*r.W).Header().Set("Access-Control-Allow-Headers", "*")
	(*r.W).Header().Set("Access-Control-Allow-Methods", "*")
	(*r.W).WriteHeader(http.StatusOK)
}

func sendJson(status int, w *http.ResponseWriter, data responseData) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	// (*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).WriteHeader(status)
	json.NewEncoder(*w).Encode(data)
}
