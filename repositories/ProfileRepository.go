package repositories

import (
	"app/models"
	"errors"
	"fmt"
)

func GetProfileById(profileId int) (models.Profile, error) {

	var profile models.Profile

	err := db().Preload("SocialNetwork").First(&profile, "profile_id = ?", profileId)
	if err.Error != nil {
		fmt.Println("error!!!")
		return profile, errors.New("SQL ERROR")
	}

	return profile, nil
}

func GetActiveProfile() (models.Profile, error) {
	var profile models.Profile

	err := db().First(&profile, "active is true")
	if err.Error != nil {
		fmt.Println("error!!!")
		return profile, errors.New("SQL ERROR")
	}

	return profile, nil
}

func GetProfilesByUserId(userId int) ([]models.Profile, error) {
	var profiles []models.Profile

	err := db().Where("user_id = ?", userId).Find(&profiles)
	if err.Error != nil {
		fmt.Println("error!!!")
		return profiles, errors.New("SQL ERROR")
	}

	return profiles, nil
}

func Create(p models.Profile) models.Profile {

	db().Create(&p)

	return p
}
