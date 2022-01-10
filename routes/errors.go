package router

import (
	"app/responses"
	"net/http"
)

func httpNotFound(w http.ResponseWriter, r *http.Request) {

	res := responses.ResponseWriter{W: &w}

	res.NotFound("HTTP_NOT_FOUND")
}

func httpNotAllowed(w http.ResponseWriter, r *http.Request) {
	res := responses.ResponseWriter{W: &w}

	res.MethodNodAllowed("METHOD_NOT_ALLOWED")
}
