package repositories

import (
	"app/models"
	"errors"
	"fmt"
)

func GetProfileExperience(profileId int) ([]models.Experience, error) {

	var experiences []models.Experience

	err := db().Where("profile_id = ?", profileId).Find(&experiences)
	if err.Error != nil {
		fmt.Println("error!!!")
		return experiences, errors.New("SQL ERROR")
	}

	return experiences, nil
}
