package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(m *mux.Router) {

	website(m)

	authentication(m)

	user(m)

	exceptionRoutes(m)
}

func exceptionRoutes(m *mux.Router) {

	m.NotFoundHandler = http.HandlerFunc(httpNotFound)

	m.MethodNotAllowedHandler = http.HandlerFunc(httpNotAllowed)
}
