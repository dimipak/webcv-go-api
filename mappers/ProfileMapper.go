package mappers

import (
	"app/models"
	"app/requests"
)

func ProfileCreateMapper(req requests.CreateProfileRequest) models.Profile {
	timestamps := models.NowFormatted()
	return models.Profile{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		CreatedAt: timestamps,
		UpdatedAt: timestamps,
	}
}
