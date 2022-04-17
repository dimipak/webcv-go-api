package models

import (
	db "app/system"
	"errors"
	"sort"
	"time"
)

type Education struct {
	EducationId int    `json:"education_id" gorm:"primarykey"`
	ProfileId   int    `json:"profile_id"`
	Title       string `json:"title"`
	Reference   string `json:"reference"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Date        string `json:"date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Educations []Education

func (e Educations) GetOne(id int) (Education, error) {

	for i, education := range e {
		if education.EducationId == id {
			return e[i], nil
		}
	}

	return Education{}, errors.New("education does not exist")
}

func (e Educations) OrderByDate() Educations {

	sort.SliceStable(e, func(i, j int) bool {
		dateI, _ := time.Parse("2006-01-02T15:04:05Z", e[i].Date)
		dateJ, _ := time.Parse("2006-01-02T15:04:05Z", e[j].Date)
		return dateI.After(dateJ)
	})
	return e
}

func (e *Education) Update(education Education) error {
	return db.GORM().Model(e).Updates(education).Error
}

func (e *Education) Delete() error {
	return db.GORM().Delete(e).Error
}

func (e Educations) Delete() error {
	return db.GORM().Delete(e).Error
}
