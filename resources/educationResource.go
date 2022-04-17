package resources

import (
	"app/models"
	"time"
)

type Education struct {
	EducationId int      `json:"education_id"`
	ProfileId   int      `json:"profile_id"`
	Title       string   `json:"title"`
	Reference   string   `json:"reference"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	Date        string   `json:"date"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

func EducationsResources(educations []models.Education) interface{} {

	var filtered []interface{}

	for _, education := range educations {

		filtered = append(filtered, EducationResources(education))
	}

	return filtered
}

func EducationResources(education models.Education) Education {

	layout := "2006-01-02T15:04:05Z"
	format := "2006-01-02 15:04:05"

	createdAt, _ := time.Parse(layout, education.CreatedAt)
	updatedAt, _ := time.Parse(layout, education.UpdatedAt)
	date, _ := time.Parse(layout, education.Date)

	return Education{
		EducationId: education.EducationId,
		ProfileId:   education.ProfileId,
		Title:       education.Title,
		Reference:   education.Reference,
// 		Description: strings.Split(education.Description, "\n"),
		Description: education.Description,
		Link:        education.Link,
		Date:        date.Format(format),
		CreatedAt:   createdAt.Format(format),
		UpdatedAt:   updatedAt.Format(format),
	}
}
