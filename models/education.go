package models

type Education struct {
	EducationId int    `json:"education_id"`
	ProfileId   int    `json:"profile_id"`
	Title       string `json:"title"`
	Reference   string `json:"reference"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Date        string `json:"date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
