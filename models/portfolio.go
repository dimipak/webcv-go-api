package models

import (
	db "app/config"
	"errors"
)

type Portfolio struct {
	PortfolioId int    `json:"portfolio_id" gorm:"primarykey"`
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

type Portfolios []Portfolio

func (p *Portfolio) First(id int) Portfolio {
	var portfolio Portfolio
	db.GORM().Where("portfolio_id = ?").First(&portfolio)

	return portfolio
}

func (p Portfolios) GetOne(id int) (Portfolio, error) {

	for i, portfolio := range p {
		if portfolio.PortfolioId == id {
			return p[i], nil
		}
	}

	return Portfolio{}, errors.New("portfolio does not exist")
}

func (p *Portfolio) Update(portfolio Portfolio) error {
	return db.GORM().Model(p).Updates(portfolio).Error
}

func (p *Portfolio) Delete() error {
	return db.GORM().Delete(p).Error
}
