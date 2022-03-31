package mappers

import (
	"app/models"
	"app/requests"
)

func UpdateProfileMapper(req requests.UpdateProfileRequest) models.Profile {
	return models.Profile{
		Username:    req.Username,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		FirstQuote:  req.FirstQuote,
		SecondQuote: req.SecondQuote,
		About:       req.About,
	}
}
