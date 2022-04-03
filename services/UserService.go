package services

import (
	"app/models"
	m "app/models"
	"app/repositories"
	"app/requests"
	"app/system"

	"fmt"
	"net/http"
	"strconv"

	"errors"

	"github.com/gorilla/mux"
)

type UserService struct {
	UserId    int
	ProfileId int
}

func (u *UserService) GetById() (m.User, error) {
	var userRepository repositories.UserRepository

	return userRepository.SetUserId(u.UserId).GetById()
}

func (u *UserService) SetUserId(id int) *UserService {
	u.UserId = id
	return u
}

func UserRegister(ur requests.UserRegisterRequest) (m.User, error) {

	user, _ := repositories.GetUserByEmail(ur.Email)

	if user.UserId == 0 || user.Activated {
		return user, errors.New("user already exist")
	}

	newUser := m.NewUser(ur.Username, ur.Email, ur.Password)

	createdUser := repositories.CreateUser(newUser)

	go system.SendMail(newUser.Email, "http://dimipak.test/activate/key/"+newUser.ActivateKey)

	return createdUser, nil
}

func ActivateUser(k string) (m.User, error) {

	user, _ := repositories.GetUserByActivateKey(k)

	if user.UserId == 0 || user.Activated {
		return user, errors.New("wrong activation key")
	}

	user.Activate()

	return user, nil
}

func Login(ul requests.UserLoginRequest) (m.User, error) {

	user, _ := repositories.GetUserByUsername(ul.Username)

	if user.UserId == 0 || !user.Activated {
		return user, errors.New("user already exist")
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

func UpdateProfileImage(profileId int, url string) m.Profile {

	profile, _ := repositories.FindProfileById(profileId)

	profile.UpdateProfileImage(url)

	return profile
}
