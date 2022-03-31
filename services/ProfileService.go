package services

import (
	"app/models"
	m "app/models"
	"app/repositories"
	"app/requests"
)

type ProfileService struct {
	ProfileId            int
	UserId               int
	Profile              m.Profile
	ProfileUpdateRequest requests.UpdateProfileRequest
}

var (
	profileRepository       repositories.ProfileRepository
	socialNetworkRepository repositories.SocialNetworkRepository
)

func (p *ProfileService) SetProfileId(profileId int) *ProfileService {
	p.ProfileId = profileId

	return p
}

func (p *ProfileService) SetProfileUpdateRequest(req requests.UpdateProfileRequest) *ProfileService {
	p.ProfileUpdateRequest = req

	return p
}

func (p *ProfileService) SetUserId(userId int) *ProfileService {
	p.UserId = userId

	return p
}

func GetActiveProfileInfo(profileId int) (models.Profile, error) {

	return repositories.GetProfileById(profileId)
}

func GetActiveProfile() (models.Profile, error) {

	return repositories.GetActiveProfile()
}

func GetActiveProfileSkills(profileId int) ([]models.Skill, error) {

	return repositories.GetProfileSkills(profileId)
}

func GetActiveProfilePortfolio(profileId int) ([]models.Portfolio, error) {

	return repositories.GetProfilePortfolio(profileId)
}

func GetActiveProfileExperiences(profileId int) ([]models.Experience, error) {

	return repositories.GetProfileExperience(profileId)
}

func GetActiveProfileEducations(profileId int) ([]models.Education, error) {

	return repositories.GetProfileEducations(profileId)
}

func (p *ProfileService) GetById() (models.Profile, error) {
	return repositories.GetProfileById(p.ProfileId)
}

func (p *ProfileService) GetUserProfile() (models.Profile, error) {
	return profileRepository.SetProfileId(p.ProfileId).SetUserId(p.UserId).GetByUserIdAndProfileId()
}

func (p *ProfileService) UpdateById(newProfile func(requests.UpdateProfileRequest) models.Profile) (models.Profile, error) {

	profileRepository.ProfileId = p.ProfileId
	socialNetworkRepository.ProfileId = p.ProfileId

	profile, err := profileRepository.UpdateById(newProfile(p.ProfileUpdateRequest))
	if err != nil {
		return profile, err
	}

	socialNetwork, err := socialNetworkRepository.GetByProfileId()
	if err != nil {
		return profile, err
	}

	socialNetworkRepository.SocialNetworkId = socialNetwork.SocialNetworkId
	socialNetwork, err = socialNetworkRepository.UpdateById(models.SocialNetwork{
		Linkedin: p.ProfileUpdateRequest.SocialNetwork.Linkedin,
		Github:   p.ProfileUpdateRequest.SocialNetwork.Github,
	})
	if err != nil {
		return profile, err
	}

	return profile, nil
}
