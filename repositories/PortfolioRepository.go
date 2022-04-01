package repositories

import (
	db "app/config"
	"app/models"
	"errors"
	"fmt"
)

type PortfolioRepository struct {
	profileId   int
	portfolioId int
	profile     models.Profile
	portfolio   models.Portfolio
}

func (p *PortfolioRepository) SetProfileId(profileId int) *PortfolioRepository {
	p.profileId = profileId
	return p
}

func (p *PortfolioRepository) SetPortfolioId(portfolioId int) *PortfolioRepository {
	p.portfolioId = portfolioId
	return p
}

func (p *PortfolioRepository) Create(newPortfolio models.Portfolio) (models.Portfolio, error) {
	newPortfolio.CreatedAt = models.NowFormatted()
	newPortfolio.UpdatedAt = models.NowFormatted()

	res := db.GORM().Create(&newPortfolio)

	return newPortfolio, res.Error
}

func GetProfilePortfolio(profileId int) ([]models.Portfolio, error) {

	var portfolio []models.Portfolio

	err := db.GORM().Where("profile_id = ?", profileId).Find(&portfolio)
	if err.Error != nil {
		fmt.Println("error!!!")
		return portfolio, errors.New("SQL ERROR")
	}

	return portfolio, nil
}
