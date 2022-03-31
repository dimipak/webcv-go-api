package services

import (
	"app/helpers"
	"app/models"
	"app/repositories"
	"app/requests"
	"errors"
	"strconv"
)

type SkillService struct {
	ProfileId              int
	SkillId                int
	Skill                  models.Skill
	Skills                 []models.Skill
	CreateSkillRequestBody requests.SkillCreateRequest
	UpdateSkillRequestBody requests.SkillUpdateRequest
}

func (r *SkillService) GetByProfileId() ([]models.Skill, error) {
	var skillRepository repositories.SkillRepository

	skillRepository.ProfileId = r.ProfileId

	return skillRepository.GetByProfileId()
}

func (r *SkillService) GetById() (models.Skill, error) {
	var skillRepository repositories.SkillRepository

	skillRepository.SkillId = r.SkillId

	return skillRepository.GetById()
}

func (r *SkillService) UpdateByIdOld(newSkill models.Skill) (models.Skill, error) {
	var skillRepository repositories.SkillRepository

	skillRepository.SkillId = r.SkillId

	return skillRepository.UpdateById(newSkill)
}

func (r *SkillService) UpdateById(newSkill func(requests.SkillUpdateRequest) models.Skill) (models.Skill, error) {
	var skillRepository repositories.SkillRepository

	skillRepository.SkillId = r.SkillId

	return skillRepository.UpdateById(newSkill(r.UpdateSkillRequestBody))
}

func (r *SkillService) Create(newSkill func(requests.SkillCreateRequest, int) models.Skill) (models.Skill, error) {
	var skillRepository repositories.SkillRepository

	return skillRepository.Create(newSkill(r.CreateSkillRequestBody, r.ProfileId))
}

func (r *SkillService) DeleteById() (models.Skill, error) {
	var skillRepository repositories.SkillRepository

	skillRepository.SkillId = r.SkillId

	return skillRepository.DeleteById()
}

func (r *SkillService) UpdateSkillsOrder(skillIds []int) ([]models.Skill, error) {
	var skillRepository repositories.SkillRepository

	skillRepository.ProfileId = r.ProfileId

	skills, err := skillRepository.GetByProfileId()
	if err != nil {
		return skills, err
	}

	for _, skill := range skills {
		if !helpers.InArray(skill.SkillId, skillIds) {
			return skills, errors.New("skill id " + strconv.Itoa(skill.SkillId) + "  does not exist")
		}
	}

	if len(skills) != len(skillIds) {
		return skills, errors.New("number of skills are wrong")
	}

	return skillRepository.UpdateSkillsOrder(skillIds)
}
