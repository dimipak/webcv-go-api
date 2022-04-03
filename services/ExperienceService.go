package services

import (
	"app/models"
	"app/repositories"
	"app/requests"
	"errors"
	"fmt"
	"time"
)

type ExperienceService struct {
	profileId               int
	experienceId            int
	experience              models.Experience
	experienceCreateRequest requests.ExperienceCreateRequest
}

var (
	experienceRepository repositories.ExperienceRepository
)

func (e *ExperienceService) SetProfileId(profileId int) *ExperienceService {
	e.profileId = profileId
	return e
}

func (e *ExperienceService) SetExperienceId(experienceId int) *ExperienceService {
	e.experienceId = experienceId
	return e
}

func (e *ExperienceService) SetCreateRequest(req requests.ExperienceCreateRequest) *ExperienceService {
	e.experienceCreateRequest = req
	return e
}

func (e *ExperienceService) SetExperience(exp models.Experience) *ExperienceService {
	e.experience = exp
	return e
}

func (e *ExperienceService) GetByProfileId() ([]models.Experience, error) {
	return experienceRepository.SetProfileId(e.profileId).GetByProfileId()
}

func (e *ExperienceService) checkExperienceDates(experience models.Experience, experiences []models.Experience) bool {

	newExpStart, _ := time.Parse("2006-01-02 15:04:05", experience.StartDate)
	newExpEnd, _ := time.Parse("2006-01-02 15:04:05", experience.EndDate)

	for _, exp := range experiences {
		start, _ := time.Parse("2006-01-02T15:04:05Z", exp.StartDate)
		if exp.EndDate == "" {
			fmt.Println("yes its empty")
		}
		end, _ := time.Parse("2006-01-02T15:04:05Z", exp.EndDate)

		if (newExpStart.After(start) && newExpStart.Before(end)) || (newExpEnd.After(start) && newExpEnd.Before(end)) {
			fmt.Println("found date conflict")
			return false
		}
	}

	return true
}

func (e *ExperienceService) Create(exp func(requests.ExperienceCreateRequest) models.Experience) (models.Experience, error) {

	experience := exp(e.experienceCreateRequest)
	experience.ProfileId = e.profileId

	experiences, err := e.GetByProfileId()
	if err != nil {
		return models.Experience{}, err
	}

	if !e.checkExperienceDates(experience, experiences) {
		return experience, errors.New("wrong dates provided")
	}

	return experienceRepository.Create(experience)
}

func (e *ExperienceService) Update(newExperience models.Experience) (models.Experience, error) {

	experiences, err := e.GetByProfileId()
	if err != nil {
		return models.Experience{}, err
	}

	if !e.checkExperienceDates(e.experience, experiences) {
		return e.experience, errors.New("wrong dates provided")
	}

	err = e.experience.Update(newExperience)
	if err != nil {
		return e.experience, err
	}

	return e.experience, nil
}
