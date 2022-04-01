package mappers

import (
	"app/models"
	"app/requests"
)

func PortfolioCreateMapper(req requests.PortfolioCreateRequest) models.Portfolio {
	return models.Portfolio{
		Name:       req.Name,
		Customer:   req.Customer,
		Technology: req.Technology,
		Type:       req.Type,
		WebsiteUrl: req.WebsiteUrl,
	}
}
