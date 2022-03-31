package repositories

import (
	db "app/config"
	"app/models"
	"errors"
	"fmt"
)

func GetProfileEducations(profileId int) ([]models.Education, error) {

	var educations []models.Education

	err := db.GORM().Where("profile_id = ?", profileId).Find(&educations)
	if err.Error != nil {
		fmt.Println("error!!!")
		return educations, errors.New("SQL ERROR")
	}

	return educations, nil
}
