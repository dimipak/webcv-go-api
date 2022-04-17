package resources

import (
	"app/models"
	"time"
)

type Experience struct {
	ExperienceId int    `json:"experience_id"`
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

func ExperiencesResources(experiences []models.Experience) interface{} {

	var filtered []interface{}

	for _, experience := range experiences {

		filtered = append(filtered, ExperienceResources(experience))
	}

	return filtered
}

func ExperienceResources(experience models.Experience) Experience {

	layout := "2006-01-02T15:04:05Z"
	format := "2006-01-02 15:04:05"

	createdAt, _ := time.Parse(layout, experience.CreatedAt)
	updatedAt, _ := time.Parse(layout, experience.UpdatedAt)

	startDate, _ := time.Parse(layout, experience.StartDate)

	var endDate string
	date, err := time.Parse(layout, experience.EndDate)
	if err != nil {
		endDate = "-"
	} else {
		endDate = date.Format(format)
	}

	return Experience{
		ExperienceId: experience.ExperienceId,
		ProfileId:    experience.ProfileId,
		CompanyName:  experience.CompanyName,
		Role:         experience.Role,
		// Description:  strings.Split(experience.Description, "\n"),
		Description: experience.Description,
		Country:     experience.Country,
		City:        experience.City,
		StartDate:   startDate.Format(format),
		EndDate:     endDate,
		CreatedAt:   createdAt.Format(format),
		UpdatedAt:   updatedAt.Format(format),
	}
}
