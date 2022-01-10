package responses

import (
	"encoding/json"
	"net/http"
)

const appJson string = "application/json"

type ResponseWriter struct {
	W *http.ResponseWriter
}

type ResponseData struct {
	Code        int         `json:"code"`
	Status      string      `json:"status"`
	AccessToken string      `json:"access_token,omitempty"`
	TokenType   string      `json:"token_type,omitempty"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	Error       bool        `json:"error"`
	ErrorType   string      `json:"error_type"`
}

func (w *ResponseWriter) Success(m string, data interface{}) {

	res := ResponseData{
		Code:    http.StatusOK,
		Status:  "success",
		Message: m,
		Data:    data,
		Error:   false,
	}

	(*w.W).Header().Set("Content-Type", appJson)
	(*w.W).WriteHeader(http.StatusOK)
	json.NewEncoder(*w.W).Encode(res)
}

func (w *ResponseWriter) BadRequest(m string) {

	res := ResponseData{
		Code:      http.StatusBadRequest,
		Status:    "error",
		Message:   m,
		Data:      make([]interface{}, 0),
		Error:     true,
		ErrorType: "bad_request",
	}

	(*w.W).Header().Set("Content-Type", appJson)
	(*w.W).WriteHeader(http.StatusBadRequest)
	json.NewEncoder(*w.W).Encode(res)
}

func (w *ResponseWriter) NotFound(m string) {

	res := ResponseData{
		Code:      http.StatusNotFound,
		Status:    "error",
		Message:   m,
		Data:      make([]interface{}, 0),
		Error:     true,
		ErrorType: "not_found",
	}

	(*w.W).Header().Set("Content-Type", appJson)
	(*w.W).WriteHeader(http.StatusNotFound)
	json.NewEncoder(*w.W).Encode(res)
}

func (w *ResponseWriter) MethodNodAllowed(m string) {

	res := ResponseData{
		Code:      http.StatusMethodNotAllowed,
		Status:    "error",
		Message:   m,
		Data:      make([]interface{}, 0),
		Error:     true,
		ErrorType: "not_allowed",
	}

	(*w.W).Header().Set("Content-Type", appJson)
	(*w.W).WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(*w.W).Encode(res)
}

func (w *ResponseWriter) SuccessWithToken(m string, token string, data interface{}) {

	res := ResponseData{
		Code:        http.StatusOK,
		Status:      "success",
		Message:     m,
		AccessToken: token,
		TokenType:   "bearer",
		Data:        data,
		Error:       false,
	}

	(*w.W).Header().Set("Content-Type", appJson)
	(*w.W).WriteHeader(http.StatusOK)
	json.NewEncoder(*w.W).Encode(res)
}

func (w *ResponseWriter) Unauthorized() {

	res := ResponseData{
		Code:      http.StatusUnauthorized,
		Status:    "error",
		Message:   "User is not authorized",
		Data:      make([]interface{}, 0),
		Error:     true,
		ErrorType: "unauthorized",
	}

	(*w.W).Header().Set("Content-Type", appJson)
	(*w.W).WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(*w.W).Encode(res)
}
