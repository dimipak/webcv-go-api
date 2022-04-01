package mappers

import (
	"app/models"
	"app/requests"
	"time"
)

func EducationCreateMapper(req requests.EducationCreateRequest) models.Education {

	date, _ := time.Parse("2006", req.Date)

	return models.Education{
		Title:       req.Title,
		Reference:   req.Reference,
		Description: req.Description,
		Link:        req.Link,
		Date:        date.Format("2006-01-02 15:04:05"),
	}
}

func EducationUpdateMapper(req requests.EducationUpdateRequest) models.Education {

	date, _ := time.Parse("2006", req.Date)

	return models.Education{
		Title:       req.Title,
		Reference:   req.Reference,
		Description: req.Description,
		Link:        req.Link,
		Date:        date.Format("2006-01-02 15:04:05"),
	}
}
