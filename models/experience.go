package models

type Experience struct {
	ExperienceId int    `json:"experience_id"`
	ProfileId    int    `json:"profile_id"`
	CompanyName  string `json:"company_name"`
	Role         string `json:"role"`
	Description  string `json:"description"`
	Country      string `json:"country"`
	City         string `json:"city"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
