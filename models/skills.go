package models

type Skill struct {
	SkillId     int    `json:"skill_id"`
	ProfileId   int    `json:"profile_id"`
	Name        string `json:"name"`
	Progress    int    `json:"progress"`
	Description string `json:"description"`
	Order       int    `json:"order"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
