package routes

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

	user.HandleFunc("/{user_id}/profiles/{profile_id}/profile-image", ctrl.UploadProfileImage).Methods("POST")

	user.HandleFunc("/{user_id}/profiles/{profile_id}", ctrl.UpdateProfile).Methods("PUT")

	user.HandleFunc("/{user_id}/profiles/{profile_id}/skills", ctrl.GetUserProfileSkills).Methods("GET")

	user.HandleFunc("/{user_id}/profiles/{profile_id}/skills/{skill_id}", ctrl.GetUserProfileSkill).Methods("GET")

	user.HandleFunc("/{user_id}/profiles/{profile_id}/skills/{skill_id}", ctrl.UpdateUserProfileSkill).Methods("PUT")

	user.HandleFunc("/{user_id}/profiles/{profile_id}/skills", ctrl.UpdateUserProfileSkillsOrder).Methods("PUT")

	user.HandleFunc("/{user_id}/profiles/{profile_id}/skills/{skill_id}", ctrl.DeleteUserProfileSkill).Methods("DELETE")
	
	user.HandleFunc("/{user_id}/profiles/{profile_id}/skills", ctrl.CreateUserProfileSkill).Methods("POST")
}
