package routes

import (
	ctrl "app/controllers"

	"github.com/gorilla/mux"
)

func authentication(m *mux.Router) {

	var userController ctrl.UserController

	m.HandleFunc("/register", userController.Register).Methods("POST")

	m.HandleFunc("/activate/key/{key}", userController.ActivateUser).Methods("GET")

	m.HandleFunc("/login", userController.Login).Methods("POST")
}
