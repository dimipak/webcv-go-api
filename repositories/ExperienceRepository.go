package repositories

import (
	db "app/config"
	"app/models"
	"errors"
	"fmt"
)

type ExperienceRepository struct {
	profileId    int
	experienceId int
	expperience  models.Experience
}

func (e *ExperienceRepository) SetProfileId(profileId int) *ExperienceRepository {
	e.profileId = profileId
	return e
}

func (e *ExperienceRepository) SetExperienceId(experienceId int) *ExperienceRepository {
	e.experienceId = experienceId
	return e
}

func (e *ExperienceRepository) SetExperience(exp models.Experience) *ExperienceRepository {
	e.expperience = exp
	return e
}

func (e *ExperienceRepository) GetByProfileId() ([]models.Experience, error) {
	var experiences []models.Experience

	res := db.GORM().Where("profile_id = ?", e.profileId).Find(&experiences)

	return experiences, res.Error
}

func (e *ExperienceRepository) Create(experience models.Experience) (models.Experience, error) {
	experience.CreatedAt = models.NowFormatted()
	experience.UpdatedAt = models.NowFormatted()

	res := db.GORM().Create(&experience)

	return experience, res.Error
}

func (e *ExperienceRepository) Update(experience models.Experience) (models.Experience, error) {
	var exp models.Experience

	res := db.GORM().Model(exp).Updates(&experience)

	return experience, res.Error
}

func GetProfileExperience(profileId int) ([]models.Experience, error) {

	var experiences []models.Experience

	err := db.GORM().Where("profile_id = ?", profileId).Find(&experiences)
	if err.Error != nil {
		fmt.Println("error!!!")
		return experiences, errors.New("SQL ERROR")
	}

	return experiences, nil
}
