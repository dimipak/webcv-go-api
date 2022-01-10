package resources

import (
	"app/models"
	"strings"
)

type SocialNetwork struct {
	Linkedin      string `json:"linkedin,omitempty"`
	Github        string `json:"github,omitempty"`
	Stackoverflow string `json:"stackoverflow,omitempty"`
}

type Profile struct {
	ProfileId     int            `json:"profile_id,omitempty"`
	UserId        int            `json:"user_id,omitempty"`
	Active        bool           `json:"active,omitempty"`
	Username      string         `json:"username,omitempty"`
	FirstName     string         `json:"first_name,omitempty"`
	LastName      string         `json:"last_name,omitempty"`
	FirstQuote    string         `json:"first_quote,omitempty"`
	SecondQuote   string         `json:"second_quote,omitempty"`
	Email         string         `json:"email,omitempty"`
	Phone         string         `json:"phone,omitempty"`
	About         []string       `json:"about,omitempty"`
	SocialNetwork *SocialNetwork `json:"social_networks,omitempty"`
	ProfileImage  string         `json:"profile_image,omitempty"`
	CoverImage    string         `json:"cover_image,omitempty"`
	CreatedAt     string         `json:"created_at,omitempty"`
	UpdatedAt     string         `json:"updated_at,omitempty"`
}

func ProfileResource(profile models.Profile) Profile {
	return Profile{
		ProfileId:     profile.ProfileId,
		UserId:        profile.UserId,
		Active:        profile.Active,
		Username:      profile.Username,
		FirstName:     profile.FirstName,
		LastName:      profile.LastName,
		FirstQuote:    profile.FirstName,
		SecondQuote:   profile.SecondQuote,
		Email:         profile.Email,
		About:         strings.Split(profile.About, "\n"),
		SocialNetwork: SocialNetworkResource(profile.SocialNetwork),
		ProfileImage:  profile.ProfileImage,
		CoverImage:    profile.CoverImage,
	}
}

func SocialNetworkResource(sn models.SocialNetwork) *SocialNetwork {
	return &SocialNetwork{
		Linkedin:      sn.Linkedin,
		Github:        sn.Github,
		Stackoverflow: sn.StackOverflow,
	}
}

func ActiveProfileResource(profile models.Profile) Profile {
	return Profile{
		ProfileId: profile.ProfileId,
		Username:  profile.Username,
		Active:    profile.Active,
	}
}
