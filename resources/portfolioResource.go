package resources

import (
	"app/models"
	"time"
)

type Portfolio struct {
	PortfolioId int    `json:"portfolio_id"`
	ProfileId   int    `json:"profile_id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Technology  string `json:"technology"`
	Customer    string `json:"customer"`
	ImageUrl    string `json:"image_url"`
	WebsiteUrl  string `json:"website_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func PortfoliosResources(portfolios []models.Portfolio) interface{} {

	var filteredPortfolios []interface{}

	for _, portfolio := range portfolios {

		filteredPortfolios = append(filteredPortfolios, PortfolioResources(portfolio))
	}

	return filteredPortfolios
}

func PortfolioResources(portfolio models.Portfolio) Portfolio {

	layout := "2006-01-02T15:04:05Z"
	format := "2006-01-02 15:04:05"

	createdAt, _ := time.Parse(layout, portfolio.CreatedAt)
	updatedAt, _ := time.Parse(layout, portfolio.UpdatedAt)

	return Portfolio{
		PortfolioId: portfolio.PortfolioId,
		ProfileId:   portfolio.ProfileId,
		Name:        portfolio.Name,
		Type:        portfolio.Type,
		Technology:  portfolio.Technology,
		Customer:    portfolio.Customer,
		ImageUrl:    portfolio.ImageUrl,
		WebsiteUrl:  portfolio.WebsiteUrl,
		CreatedAt:   createdAt.Format(format),
		UpdatedAt:   updatedAt.Format(format),
	}
}
