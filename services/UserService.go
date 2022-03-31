package services

import (
	"app/models"
	m "app/models"
	"app/repositories"
	"app/requests"
	"app/systemService"
	"fmt"
	"net/http"
	"strconv"

	"app/systemService/email"

	"errors"

	"github.com/gorilla/mux"
)

// type UserServiceInterface interface {
// 	UpdateUserProfile(profileId int, newProfile m.Profile, newSocialNetworks m.SocialNetwork) m.Profile
// 	GetUserProfileSkills() ([]m.Skill, error)
// }

type UserService struct {
	ProfileId int
}

func UserRegister(ur requests.UserRegisterRequest) (m.User, error) {

	user, _ := repositories.GetUserByEmail(ur.Email)

	if user != (m.User{}) || user.Activated {
		return user, errors.New("user already exist")
	}

	newUser := m.NewUser(ur.Username, ur.Email, ur.Password)

	createdUser := repositories.CreateUser(newUser)

	go email.SendMail(newUser.Email, "http://dimipak.test/activate/key/"+newUser.ActivateKey)

	return createdUser, nil
}

func ActivateUser(k string) (m.User, error) {

	user, _ := repositories.GetUserByActivateKey(k)

	if user == (m.User{}) || user.Activated {
		return user, errors.New("wrong activation key")
	}

	user.Activate()

	return user, nil
}

func Login(ul requests.UserLoginRequest) (m.User, error) {

	user, _ := repositories.GetUserByUsername(ul.Username)

	if user == (m.User{}) || !user.Activated {
		return user, errors.New("user already exist")
	}

	if !systemService.ComparePasswords(user.Password, ul.Password) {
		return user, errors.New("wrong passowrd")
	}

	return user, nil
}

func GetUserProfiles(userId int) ([]m.Profile, error) {
	return repositories.GetProfilesByUserId(userId)
}

func GetUserProfile(userId int, profileId int) (m.Profile, error) {
	profile, err := repositories.GetProfileById(profileId)
	if err != nil {
		return profile, err
	}

	if profile.UserId != userId {
		return models.Profile{}, errors.New("wrong profile id provided")
	}

	return profile, nil
}

func ActivateUserProfile(r *http.Request) (m.Profile, error) {
	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		fmt.Println("url vars error")
		return m.Profile{}, err
	}

	profileId, err := strconv.Atoi(vars["profile_id"])
	if err != nil {
		fmt.Println("url vars error")
		return m.Profile{}, err
	}

	fmt.Println("profile id = ", profileId)

	profiles, err := repositories.GetProfilesByUserId(userId)
	if err != nil {
		fmt.Println("profiles retrieve error")
		return m.Profile{}, err
	}

	return activateProfile(profiles, profileId)
}

func activateProfile(profiles []m.Profile, profileId int) (m.Profile, error) {

	var activeProfile m.Profile

	var targetProfile m.Profile

	for _, profile := range profiles {
		if profile.ProfileId == profileId {
			targetProfile = profile
		}
		if profile.Active {
			activeProfile = profile
		}
	}

	if targetProfile.ProfileId == 0 {
		return targetProfile, errors.New("profile does not exist")
	}

	fmt.Println("target profile id = ", targetProfile.ProfileId)
	fmt.Println("active profile id = ", activeProfile.ProfileId)

	activeProfile.DeActivate()

	targetProfile.Activate()

	return targetProfile, nil
}

func CreateProfile(userId int, request requests.CreateProfileRequest) m.Profile {
	newProfile := m.NewProfile(userId, request.Username, request.FirstName, request.LastName)

	createdProfile := repositories.Create(newProfile)

	return createdProfile
}

func UpdateProfileImage(profileId int, url string) m.Profile {

	profile, _ := repositories.FindProfileById(profileId)

	profile.UpdateProfileImage(url)

	return profile
}

// func UpdateUserProfile(profileId int, newProfile m.Profile, newSocialNetworks m.SocialNetwork) m.Profile {

// 	profile := repositories.GetProfile()
// 	socialNetworks := repositories.GetSN()

// 	repos := repositories.Repositories{
// 		ProfileRepository:       &profile,
// 		SocialNetworkRepository: &socialNetworks,
// 	}

// 	err := repos.ProfileRepository.GetById(profileId)
// 	if err != nil {
// 		fmt.Println("No profile found")
// 	}

// 	repos.ProfileRepository.Update(newProfile)

// 	repos.SocialNetworkRepository.GetByProfileId(profileId)

// 	repos.SocialNetworkRepository.Update(newSocialNetworks)

// 	repos.ProfileRepository.Preload("SocialNetwork")

// 	return profile.Profile
// }

// func GetUserProfileSkills(profileId int) ([]m.Skill, error) {
// 	skillRepository := repositories.SkillRepositoryData{
// 		ProfileId: profileId,
// 	}

// 	return skillRepository.GetByProfileId()
// }

func (us *UserService) GetUserProfileSkills() ([]m.Skill, error) {
	skillRepository := repositories.SkillRepository{
		ProfileId: us.ProfileId,
	}

	return skillRepository.GetByProfileId()
}
