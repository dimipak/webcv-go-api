package repositories

import (
	"app/models"
	db "app/system"
	"errors"
	"fmt"
)

type SkillRepositoryInterface interface {
	GetByProfileId(profileId int) ([]models.Skill, error)
}

type SkillRepository struct {
	ProfileId int
	SkillId   int
	Skill     models.Skill
	Skills    []models.Skill
}

func GetProfileSkills(profileId int) ([]models.Skill, error) {

	var skills []models.Skill

	err := db.GORM().Where("profile_id = ?", profileId).Find(&skills)
	if err.Error != nil {
		fmt.Println("error!!!")
		return skills, errors.New("SQL ERROR")
	}

	return skills, nil
}

func (r *SkillRepository) GetById() (models.Skill, error) {
	var skill models.Skill

	// err := db.GORM().First(&skills, "profile_id = ?", profileId)
	err := db.GORM().Where("skill_id = ?", r.SkillId).First(&skill)
	if err.Error != nil {
		fmt.Println("error!!!")
		return skill, errors.New("SQL ERROR")
	}

	return skill, nil
}

func (r *SkillRepository) GetByProfileId() ([]models.Skill, error) {
	var skills []models.Skill

	// err := db.GORM().First(&skills, "profile_id = ?", profileId)
	err := db.GORM().Where("profile_id = ?", r.ProfileId).Order("`order`").Find(&skills)
	if err.Error != nil {
		fmt.Println("error!!!")
		return skills, errors.New("SQL ERROR")
	}

	return skills, nil
}

func (r *SkillRepository) UpdateById(newSkill models.Skill) (models.Skill, error) {

	skill, err := r.GetById()
	if err != nil {
		return skill, err
	}

	res := db.GORM().Model(&skill).Updates(newSkill)

	return skill, res.Error
}

func (r *SkillRepository) Create(newSkill models.Skill) (models.Skill, error) {

	newSkill.CreatedAt = models.NowFormatted()
	newSkill.UpdatedAt = models.NowFormatted()

	var skill models.Skill

	err := db.GORM().Where("profile_id = ?", newSkill.ProfileId).Order("`order` desc").First(&skill)
	if err.Error != nil {
		fmt.Println("error!!!")
		return newSkill, errors.New("SQL ERROR")
	}

	newSkill.Order = skill.Order + 1

	res := db.GORM().Create(&newSkill)

	return newSkill, res.Error
}

func (r *SkillRepository) DeleteById() (models.Skill, error) {

	skill, err := r.GetById()
	if err != nil {
		return skill, err
	}

	res := db.GORM().Delete(&skill)

	return skill, res.Error
}

func (r *SkillRepository) UpdateSkillsOrder(skillIds []int) ([]models.Skill, error) {

	var skills []models.Skill

	skills, err := r.GetByProfileId()
	if err != nil {
		return skills, err
	}

	for i, skill := range skills {
		if skill.SkillId != skillIds[i] {
			for k, skillId := range skillIds {
				if skill.SkillId == skillId {
					res := db.GORM().Model(&skill).Updates(models.Skill{
						Order:     k + 1,
						UpdatedAt: models.NowFormatted(),
					})
					if res.Error != nil {
						return skills, res.Error
					}
				}
			}
		}
	}

	return r.GetByProfileId()
}
