package routes

import (
	ctrl "app/controllers"
	"app/middleware"

	"github.com/gorilla/mux"
)

func user(m *mux.Router) {

	user := m.PathPrefix("/user").Subrouter()

	user.Use(middleware.Authentication)

	profile(user)

	skills(user)

	portfolio(user)

	experience(user)

	education(user)
}

func profile(profile *mux.Router) {

	var profileController ctrl.ProfileController

	// Get all profiles
	profile.HandleFunc("/{user_id}/profiles", profileController.Get).Methods("GET")

	// Get single profile
	profile.HandleFunc("/{user_id}/profiles/{profile_id}", profileController.Show).Methods("GET")

	// Activate an inactive profile
	profile.HandleFunc("/{user_id}/profiles/{profile_id}/activate", profileController.Activate).Methods("PUT")

	// Update a profile
	profile.HandleFunc("/{user_id}/profiles/{profile_id}", profileController.Update).Methods("PUT")

	// Creates new profile
	profile.HandleFunc("/{user_id}/profiles/create", profileController.Create).Methods("POST")

	// Upload profile image
	profile.HandleFunc("/{user_id}/profiles/{profile_id}/profile-image", profileController.UploadProfileImage).Methods("POST")

	// Upload cover image
	profile.HandleFunc("/{user_id}/profiles/{profile_id}/cover-image", profileController.UploadCoverImage).Methods("POST")
}

func skills(skills *mux.Router) {

	var skillController ctrl.SkillController

	skills.HandleFunc("/{user_id}/profiles/{profile_id}/skills", skillController.Get).Methods("GET")

	skills.HandleFunc("/{user_id}/profiles/{profile_id}/skills/{skill_id}", skillController.Show).Methods("GET")

	skills.HandleFunc("/{user_id}/profiles/{profile_id}/skills", skillController.UpdateOrder).Methods("PUT")

	skills.HandleFunc("/{user_id}/profiles/{profile_id}/skills/{skill_id}", skillController.Update).Methods("PUT")

	skills.HandleFunc("/{user_id}/profiles/{profile_id}/skills", skillController.Create).Methods("POST")

	skills.HandleFunc("/{user_id}/profiles/{profile_id}/skills/{skill_id}", skillController.Delete).Methods("DELETE")
}

func portfolio(portfolio *mux.Router) {

	var portfolioController ctrl.PortfolioController

	portfolio.HandleFunc("/{user_id}/profiles/{profile_id}/portfolio", portfolioController.Get).Methods("GET")

	portfolio.HandleFunc("/{user_id}/profiles/{profile_id}/portfolio/{portfolio_id}", portfolioController.Show).Methods("GET")

	portfolio.HandleFunc("/{user_id}/profiles/{profile_id}/portfolio/{portfolio_id}", portfolioController.Update).Methods("PUT")

	portfolio.HandleFunc("/{user_id}/profiles/{profile_id}/portfolio", portfolioController.Create).Methods("POST")

	portfolio.HandleFunc("/{user_id}/profiles/{profile_id}/portfolio/{portfolio_id}/image", portfolioController.UploadImage).Methods("POST")

	portfolio.HandleFunc("/{user_id}/profiles/{profile_id}/portfolio/{portfolio_id}", portfolioController.Delete).Methods("DELETE")
}

func experience(experience *mux.Router) {

	experienceController := ctrl.ExperienceController{}

	experience.HandleFunc("/{user_id}/profiles/{profile_id}/experiences", experienceController.Get).Methods("GET")

	experience.HandleFunc("/{user_id}/profiles/{profile_id}/experiences/{experience_id}", experienceController.Show).Methods("GET")

	experience.HandleFunc("/{user_id}/profiles/{profile_id}/experiences/{experience_id}", experienceController.Update).Methods("PUT")

	experience.HandleFunc("/{user_id}/profiles/{profile_id}/experiences", experienceController.Create).Methods("POST")

	experience.HandleFunc("/{user_id}/profiles/{profile_id}/experiences/{experience_id}", experienceController.Delete).Methods("DELETE")
}

func education(education *mux.Router) {

	var educationController ctrl.EducationController

	education.HandleFunc("/{user_id}/profiles/{profile_id}/educations", educationController.Get).Methods("GET")

	education.HandleFunc("/{user_id}/profiles/{profile_id}/educations/{education_id}", educationController.Show).Methods("GET")

	education.HandleFunc("/{user_id}/profiles/{profile_id}/educations/{education_id}", educationController.Update).Methods("PUT")

	education.HandleFunc("/{user_id}/profiles/{profile_id}/educations", educationController.Create).Methods("POST")

	education.HandleFunc("/{user_id}/profiles/{profile_id}/educations/{education_id}", educationController.Delete).Methods("DELETE")
}
