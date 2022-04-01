package repositories

import (
	db "app/config"
	"app/models"
	"errors"
	"fmt"
)

type ProfileRepository struct {
	ProfileId int
	UserId    int
	Profile   models.Profile
}

func (p *ProfileRepository) SetProfileId(profileId int) *ProfileRepository {
	p.ProfileId = profileId
	return p
}

func (p *ProfileRepository) SetUserId(userId int) *ProfileRepository {
	p.UserId = userId
	return p
}

func (p *ProfileRepository) Get() (models.Profile, error) {
	var profile models.Profile

	err := db.GORM().Where("profile_id = ? and user_id = ?", p.ProfileId, p.UserId).First(&profile)
	if err.Error != nil {
		return profile, err.Error
	}

	return profile, nil
}

func (p *ProfileRepository) GetById() (models.Profile, error) {
	var profile models.Profile

	err := db.GORM().First(&profile, "profile_id = ?", p.ProfileId)
	if err != nil {
		return profile, err.Error
	}

	return profile, nil
}

func (p *ProfileRepository) UpdateById(newProfile models.Profile) (models.Profile, error) {

	profile, err := p.GetById()
	if err != nil {
		return profile, err
	}

	res := db.GORM().Model(&profile).Updates(newProfile)

	return profile, res.Error
}

func GetProfileById(profileId int) (models.Profile, error) {

	var profile models.Profile

	err := db.GORM().Preload("SocialNetwork").First(&profile, "profile_id = ?", profileId)
	if err.Error != nil {
		fmt.Println("error!!!")
		return profile, errors.New("SQL ERROR")
	}

	return profile, nil
}

func GetActiveProfile() (models.Profile, error) {
	var profile models.Profile

	err := db.GORM().First(&profile, "active is true")
	if err.Error != nil {
		fmt.Println("error!!!")
		return profile, errors.New("SQL ERROR")
	}

	return profile, nil
}

func GetProfilesByUserId(userId int) ([]models.Profile, error) {
	var profiles []models.Profile

	err := db.GORM().Where("user_id = ?", userId).Find(&profiles)
	if err.Error != nil {
		fmt.Println("error!!!")
		return profiles, errors.New("SQL ERROR")
	}

	return profiles, nil
}

func Create(p models.Profile) models.Profile {

	db.GORM().Create(&p)

	return p
}

func FindProfileById(profileId int) (models.Profile, error) {
	var profile models.Profile

	err := db.GORM().Where("profile_id = ?", profileId).First(&profile)
	if err.Error != nil {
		fmt.Println("error!!!")
		return profile, errors.New("SQL ERROR")
	}

	return profile, nil
}
