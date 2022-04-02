package repositories

import (
	"app/models"
	db "app/system"
	"errors"
	"fmt"
)

type EducationRepository struct {
	profileId   int
	educationId int
}

func (e *EducationRepository) SetProfileId(profileId int) *EducationRepository {
	e.profileId = profileId
	return e
}

func (e *EducationRepository) SetExperienceId(experienceId int) *EducationRepository {
	e.educationId = experienceId
	return e
}

func (e *EducationRepository) Create(education models.Education) (models.Education, error) {
	education.CreatedAt = models.NowFormatted()
	education.UpdatedAt = models.NowFormatted()

	res := db.GORM().Create(&education)

	return education, res.Error
}

func GetProfileEducations(profileId int) ([]models.Education, error) {

	var educations []models.Education

	err := db.GORM().Where("profile_id = ?", profileId).Find(&educations)
	if err.Error != nil {
		fmt.Println("error!!!")
		return educations, errors.New("SQL ERROR")
	}

	return educations, nil
}
