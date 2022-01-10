package repositories

import (
	"app/models"
	"errors"
	"fmt"
)

func GetProfileSkills(profileId int) ([]models.Skill, error) {

	var skills []models.Skill

	// err := db().First(&skills, "profile_id = ?", profileId)
	err := db().Where("profile_id = ?", profileId).Find(&skills)
	if err.Error != nil {
		fmt.Println("error!!!")
		return skills, errors.New("SQL ERROR")
	}

	return skills, nil
}
