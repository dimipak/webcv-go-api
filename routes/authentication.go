package router

import (
	ctrl "app/controllers"

	"github.com/gorilla/mux"
)

func authentication(m *mux.Router) {

	m.HandleFunc("/register", ctrl.Register).Methods("POST")

	m.HandleFunc("/activate/key/{key}", ctrl.ActivateUser).Methods("GET")

	m.HandleFunc("/login", ctrl.Login).Methods("POST")
}
