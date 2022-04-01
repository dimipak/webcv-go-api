package mappers

import (
	"app/models"
	"app/requests"
)

func ExperienceCreateMapper(req requests.ExperienceCreateRequest) models.Experience {
	return models.Experience{
		CompanyName: req.CompanyName,
		Role:        req.Role,
		Description: req.Description,
		Country:     req.Country,
		City:        req.City,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}
}
