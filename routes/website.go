package routes

import (
	ctrl "app/controllers"

	"github.com/gorilla/mux"
)

func website(m *mux.Router) {

	var profileController ctrl.ProfileController

	m.HandleFunc("/profile/active", profileController.GetActiveProfile).Methods("GET")

	m.HandleFunc("/profile/{profile_id}/info", profileController.GetActiveProfileInfo).Methods("GET")

	m.HandleFunc("/profile/{profile_id}/skills", profileController.GetActiveProfileSkills).Methods("GET")

	m.HandleFunc("/profile/{profile_id}/portfolio", profileController.GetActiveProfilePortfolio).Methods("GET")

	m.HandleFunc("/profile/{profile_id}/experiences", profileController.GetActiveProfileExperiences).Methods("GET")

	m.HandleFunc("/profile/{profile_id}/educations", profileController.GetActiveProfileEducations).Methods("GET")

	m.HandleFunc("/pdf", profileController.CreatePdf).Methods("GET")
}
