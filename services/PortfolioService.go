package services

import (
	"app/models"
	"app/repositories"
	"app/requests"
)

type PortfolioService struct {
	profileId              int
	portfolioId            int
	portfolioCreateRequest requests.PortfolioCreateRequest
	profile                models.Profile
	portfolio              models.Portfolio
}

func (p *PortfolioService) SetProfileId(profileId int) *PortfolioService {
	p.profileId = profileId
	return p
}

func (p *PortfolioService) SetPortfolioId(portfolioId int) *PortfolioService {
	p.portfolioId = portfolioId
	return p
}

func (p *PortfolioService) SetPortfolioCreateRequest(req requests.PortfolioCreateRequest) *PortfolioService {
	p.portfolioCreateRequest = req
	return p
}

func (r *PortfolioService) Create(newPortfolio func(requests.PortfolioCreateRequest) models.Portfolio) (models.Portfolio, error) {
	var portfolioRepository repositories.PortfolioRepository

	portfolio := newPortfolio(r.portfolioCreateRequest)
	portfolio.ProfileId = r.profileId

	return portfolioRepository.Create(portfolio)
}
