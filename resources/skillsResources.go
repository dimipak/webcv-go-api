package resources

import (
	"app/models"
	"time"
)

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

func SkillsResources(skills []models.Skill) interface{} {

	var filteredSkills []interface{}

	for _, skill := range skills {

		filteredSkills = append(filteredSkills, SkillResources(skill))
	}

	return filteredSkills
}

func SkillResources(skill models.Skill) Skill {

	layout := "2006-01-02T15:04:05Z"
	format := "2006-01-02 15:04:05"

	createdAt, _ := time.Parse(layout, skill.CreatedAt)
	updatedAt, _ := time.Parse(layout, skill.UpdatedAt)

	return Skill{
		SkillId:     skill.SkillId,
		ProfileId:   skill.ProfileId,
		Name:        skill.Name,
		Progress:    skill.Progress,
		Description: skill.Description,
		Order:       skill.Order,
		CreatedAt:   createdAt.Format(format),
		UpdatedAt:   updatedAt.Format(format),
	}
}
