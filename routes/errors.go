package routes

import (
	res "app/responses"
	"net/http"
)

func httpNotFound(w http.ResponseWriter, r *http.Request) {

	res.JsonResponse(&w, res.NotFoundResponse{
		Message: "HTTP_NOT_FOUND",
	})
}

func httpNotAllowed(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		res.OptionsResponse(&w)
		return
	}

	res.JsonResponse(&w, res.MethodNotAllowedResponse{
		Message: "METHOD_NOT_ALLOWED",
	})
}
