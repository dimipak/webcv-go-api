package repositories

import (
	db "app/config"
	"app/models"
	"errors"
	"fmt"
)

func GetProfileExperience(profileId int) ([]models.Experience, error) {

	var experiences []models.Experience

	err := db.GORM().Where("profile_id = ?", profileId).Find(&experiences)
	if err.Error != nil {
		fmt.Println("error!!!")
		return experiences, errors.New("SQL ERROR")
	}

	return experiences, nil
}
