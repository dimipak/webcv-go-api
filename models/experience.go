package models

import (
	db "app/system"
	"errors"
	"sort"
	"time"
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

func (e Experiences) OrderByStartDate() Experiences {

	sort.SliceStable(e, func(i, j int) bool {
		startDateI, _ := time.Parse("2006-01-02T15:04:05Z", e[i].StartDate)
		startDateJ, _ := time.Parse("2006-01-02T15:04:05Z", e[j].StartDate)
		return startDateI.After(startDateJ)
		// return e[i].StartDate < e[j].StartDate
	})

	return e
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
