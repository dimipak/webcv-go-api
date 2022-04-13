package services

import (
	"app/models"
	"app/repositories"
	"app/requests"
)

type EducationService struct {
	profileId     int
	educationId   int
	education     models.Education
	createRequest requests.EducationCreateRequest
}

func (e *EducationService) SetProfileId(profileId int) *EducationService {
	e.profileId = profileId
	return e
}

func (e *EducationService) SetEducationId(educationId int) *EducationService {
	e.educationId = educationId
	return e
}

func (e *EducationService) SetCreateRequest(req requests.EducationCreateRequest) *EducationService {
	e.createRequest = req
	return e
}

func (e *EducationService) Create(set func(requests.EducationCreateRequest) models.Education) (models.Education, error) {
	var educationRepository repositories.EducationRepository

	education := set(e.createRequest)
	education.ProfileId = e.profileId

	return educationRepository.Create(education)
}
