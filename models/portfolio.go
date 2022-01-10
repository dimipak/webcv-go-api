package models

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
