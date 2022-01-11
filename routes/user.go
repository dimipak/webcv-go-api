package router

import (
	ctrl "app/controllers"
	"app/middleware"

	"github.com/gorilla/mux"
)

func user(m *mux.Router) {

	user := m.PathPrefix("/user").Subrouter()

	user.Use(middleware.Authentication)

	user.HandleFunc("/{user_id}/profiles", ctrl.GetUserProfiles).Methods("GET")

	user.HandleFunc("/{user_id}/profiles/{profile_id}", ctrl.GetUserProfile).Methods("GET")

	user.HandleFunc("/{user_id}/profiles/{profile_id}/activate", ctrl.ActivateProfile).Methods("PUT")

	user.HandleFunc("/{user_id}/profiles/create", ctrl.CreateProfile).Methods("POST")
}
