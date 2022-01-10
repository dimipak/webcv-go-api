package resources

import (
	"app/models"
)

type Profiles struct {
	ProfileId    int    `json:"profile_id"`
	Actiive      bool   `json:"active"`
	Username     string `json:"username"`
	ProfileImage string `json:"profile_image"`
}

func UserProfilesResource(profiles []models.Profile) []Profiles {

	var profilesArray []Profiles

	for _, profile := range profiles {

		profilesArray = append(profilesArray, UserProfileResource(profile))
	}

	return profilesArray
}

func UserProfileResource(profile models.Profile) Profiles {

	return Profiles{
		ProfileId:    profile.ProfileId,
		Actiive:      profile.Active,
		Username:     profile.Username,
		ProfileImage: profile.ProfileImage,
	}
}
