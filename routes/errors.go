package router

import (
	res "app/responses"
	"net/http"
)

func httpNotFound(w http.ResponseWriter, r *http.Request) {

	res.JsonResponse(res.NotFoundResponse{
		W:       &w,
		Message: "HTTP_NOT_FOUND",
	})
}

func httpNotAllowed(w http.ResponseWriter, r *http.Request) {

	res.JsonResponse(res.MethodNotAllowedResponse{
		W:       &w,
		Message: "METHOD_NOT_ALLOWED",
	})
}
