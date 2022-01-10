package services

import (
	"app/models"
	"app/repositories"
)

func GetActiveProfileInfo(profileId int) (models.Profile, error) {

	return repositories.GetProfileById(profileId)
}

func GetActiveProfile() (models.Profile, error) {

	return repositories.GetActiveProfile()
}

func GetActiveProfileSkills(profileId int) ([]models.Skill, error) {

	return repositories.GetProfileSkills(profileId)
}

func GetActiveProfilePortfolio(profileId int) ([]models.Portfolio, error) {

	return repositories.GetProfilePortfolio(profileId)
}

func GetActiveProfileExperiences(profileId int) ([]models.Experience, error) {

	return repositories.GetProfileExperience(profileId)
}

func GetActiveProfileEducations(profileId int) ([]models.Education, error) {

	return repositories.GetProfileEducations(profileId)
}
