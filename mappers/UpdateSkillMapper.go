package mappers

import (
	"app/models"
	"app/requests"
)

func UpdateSkillMapper(rsBody requests.SkillUpdateRequest) models.Skill {
	return models.Skill{
		Name:        rsBody.Name,
		Description: rsBody.Description,
		Progress:    rsBody.Progress,
	}
}

func CreateSkillMapper(rsBody requests.SkillCreateRequest, profileId int) models.Skill {
	return models.Skill{
		ProfileId:   profileId,
		Name:        rsBody.Name,
		Description: rsBody.Description,
		Progress:    rsBody.Progress,
	}
}
