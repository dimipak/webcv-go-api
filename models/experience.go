package models

import (
	db "app/system"
	"errors"
)

type Experience struct {
	ExperienceId int    `json:"experience_id" gorm:"primarykey"`
	ProfileId    int    `json:"profile_id"`
	CompanyName  string `json:"company_name"`
	Role         string `json:"role"`
	Description  string `json:"description"`
	Country      string `json:"country"`
	City         string `json:"city"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type Experiences []Experience

func (e Experiences) GetExperience(id int) (Experience, error) {

	for i, experience := range e {
		if experience.ExperienceId == id {
			return e[i], nil
		}
	}

	return Experience{}, errors.New("experience does not exist")
}

func (e *Experience) Update(experience Experience) error {
	return db.GORM().Model(e).Updates(experience).Error
}

func (e *Experience) Delete() error {
	return db.GORM().Delete(e).Error
}

func (e Experiences) Delete() error {
	return db.GORM().Delete(e).Error
}
