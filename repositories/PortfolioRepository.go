package repositories

import (
	db "app/config"
	"app/models"
	"errors"
	"fmt"
)

func GetProfilePortfolio(profileId int) ([]models.Portfolio, error) {

	var portfolio []models.Portfolio

	err := db.GORM().Where("profile_id = ?", profileId).Find(&portfolio)
	if err.Error != nil {
		fmt.Println("error!!!")
		return portfolio, errors.New("SQL ERROR")
	}

	return portfolio, nil
}
