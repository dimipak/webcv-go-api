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

	// Get all profiles
	profile.HandleFunc("/{user_id}/profiles", ctrl.GetUserProfiles).Methods("GET")

	// Get single profile
	profile.HandleFunc("/{user_id}/profiles/{profile_id}", ctrl.GetUserProfile).Methods("GET")

	// Activate an inactive profile
	profile.HandleFunc("/{user_id}/profiles/{profile_id}/activate", ctrl.ActivateProfile).Methods("PUT")

	// Update a profile
	profile.HandleFunc("/{user_id}/profiles/{profile_id}", ctrl.UpdateProfile).Methods("PUT")

	// Creates new profile
	profile.HandleFunc("/{user_id}/profiles/create", ctrl.CreateProfile).Methods("POST")

	// Upload profile image
	profile.HandleFunc("/{user_id}/profiles/{profile_id}/profile-image", ctrl.UploadProfileImage).Methods("POST")

	// Upload cover image
	profile.HandleFunc("/{user_id}/profiles/{profile_id}/cover-image", ctrl.UploadCoverImage).Methods("POST")
}

func skills(skills *mux.Router) {

	skills.HandleFunc("/{user_id}/profiles/{profile_id}/skills", ctrl.GetUserProfileSkills).Methods("GET")

	skills.HandleFunc("/{user_id}/profiles/{profile_id}/skills/{skill_id}", ctrl.GetUserProfileSkill).Methods("GET")

	skills.HandleFunc("/{user_id}/profiles/{profile_id}/skills", ctrl.UpdateUserProfileSkillsOrder).Methods("PUT")

	skills.HandleFunc("/{user_id}/profiles/{profile_id}/skills/{skill_id}", ctrl.UpdateUserProfileSkill).Methods("PUT")

	skills.HandleFunc("/{user_id}/profiles/{profile_id}/skills", ctrl.CreateUserProfileSkill).Methods("POST")

	skills.HandleFunc("/{user_id}/profiles/{profile_id}/skills/{skill_id}", ctrl.DeleteUserProfileSkill).Methods("DELETE")
}

func portfolio(portfolio *mux.Router) {

	portfolio.HandleFunc("/{user_id}/profiles/{profile_id}/portfolio", ctrl.GetUserProfilePortfolios).Methods("GET")

	portfolio.HandleFunc("/{user_id}/profiles/{profile_id}/portfolio/{portfolio_id}", ctrl.GetUserProfilePortfolio).Methods("GET")

	portfolio.HandleFunc("/{user_id}/profiles/{profile_id}/portfolio/{portfolio_id}", ctrl.UpdateUserPorfilePortfolio).Methods("PUT")

	portfolio.HandleFunc("/{user_id}/profiles/{profile_id}/portfolio", ctrl.CreateUserProfilePortfolio).Methods("POST")

	portfolio.HandleFunc("/{user_id}/profiles/{profile_id}/portfolio/{portfolio_id}/image", ctrl.UploadPortfolioImage).Methods("POST")

	portfolio.HandleFunc("/{user_id}/profiles/{profile_id}/portfolio/{portfolio_id}", ctrl.DeleteUserProfilePortfolio).Methods("DELETE")
}

func experience(experience *mux.Router) {

	experience.HandleFunc("/{user_id}/profiles/{profile_id}/experiences", ctrl.GetUserProfileExperiences).Methods("GET")

	experience.HandleFunc("/{user_id}/profiles/{profile_id}/experiences/{experience_id}", ctrl.GetUserProfileExperience).Methods("GET")

	experience.HandleFunc("/{user_id}/profiles/{profile_id}/experiences/{experience_id}", ctrl.UpdateUserProfileExperience).Methods("PUT")

	experience.HandleFunc("/{user_id}/profiles/{profile_id}/experiences", ctrl.CreateUserProfileExperience).Methods("POST")

	experience.HandleFunc("/{user_id}/profiles/{profile_id}/experiences/{experience_id}", ctrl.DeleteUserProfileExperience).Methods("DELETE")
}

func education(education *mux.Router) {

	education.HandleFunc("/{user_id}/profiles/{profile_id}/educations", ctrl.GetUserProfileEducations).Methods("GET")

	education.HandleFunc("/{user_id}/profiles/{profile_id}/educations/{education_id}", ctrl.GetUserProfileEducation).Methods("GET")

	education.HandleFunc("/{user_id}/profiles/{profile_id}/educations/{education_id}", ctrl.UpdateUserProfileEducation).Methods("PUT")

	education.HandleFunc("/{user_id}/profiles/{profile_id}/educations", ctrl.CreateUserProfileEducation).Methods("POST")

	education.HandleFunc("/{user_id}/profiles/{profile_id}/educations/{education_id}", ctrl.DeleteUserProfileEducation).Methods("DELETE")
}
