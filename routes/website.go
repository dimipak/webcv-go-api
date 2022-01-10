package router

import (
	c "app/controllers"

	"github.com/gorilla/mux"
)

func website(m *mux.Router) {

	m.HandleFunc("/profile/active", c.GetActiveProfile).Methods("GET")

	m.HandleFunc("/profile/{profile_id}/info", c.GetActiveProfileInfo).Methods("GET")

	m.HandleFunc("/profile/{profile_id}/skills", c.GetActiveProfileSkills).Methods("GET")

	m.HandleFunc("/profile/{profile_id}/portfolio", c.GetActiveProfilePortfolio).Methods("GET")

	m.HandleFunc("/profile/{profile_id}/experiences", c.GetActiveProfileExperiences).Methods("GET")

	m.HandleFunc("/profile/{profile_id}/educations", c.GetActiveProfileEducations).Methods("GET")

	m.HandleFunc("/pdf", c.CreatePdf).Methods("GET")
}
