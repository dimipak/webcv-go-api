package models

import (
	db "app/config"
	"time"
)

type profileMethods interface {
	Update(profile Profile)
}

type Profile struct {
	ProfileId     int           `json:"profile_id" gorm:"primarykey"`
	UserId        int           `json:"user_id"`
	Active        bool          `json:"active"`
	Username      string        `json:"username"`
	FirstName     string        `json:"first_name"`
	LastName      string        `json:"last_name"`
	FirstQuote    string        `json:"first_quote"`
	SecondQuote   string        `json:"second_quote"`
	Email         string        `json:"email"`
	Phone         string        `json:"phone"`
	About         string        `json:"about"`
	ProfileImage  string        `json:"profile_image" gorm:"default:'https://webcv-files.s3.eu-central-1.amazonaws.com/images/default_profile.png'"`
	CoverImage    string        `json:"cover_image" gorm:"default:'https://webcv-files.s3.eu-central-1.amazonaws.com/images/default_cover.jpg'"`
	SocialNetwork SocialNetwork `gorm:"foreignKey:ProfileId;references:ProfileId"`
	Skills        []Skill       `gorm:"foreignKey:ProfileId;references:ProfileId"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
}

func NowFormatted() string {
	return time.Now().Format(timeFormat)
}

func (p *Profile) Activate() {

	db.GORM().Model(p).Updates(Profile{
		Active:    true,
		UpdatedAt: NowFormatted(),
	})
}

func (p *Profile) DeActivate() {

	db.GORM().Model(p).Updates(map[string]interface{}{
		"active":     false,
		"updated_at": NowFormatted(),
	})
}

func NewProfile(userId int, username string, firstName string, lastName string) Profile {

	return Profile{
		UserId:    userId,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: NowFormatted(),
		UpdatedAt: NowFormatted(),
	}
}

func (p *Profile) UpdateProfileImage(url string) {

	db.GORM().Model(p).Updates(Profile{
		ProfileImage: url,
		UpdatedAt:    NowFormatted(),
	})
}

func (p *Profile) UpdateProfileCover(url string) {
	db.GORM().Model(p).Updates(Profile{
		CoverImage: url,
		UpdatedAt:  NowFormatted(),
	})
}

func (p *Profile) Update(profile Profile) {
	db.GORM().Model(p).Updates(profile)
}

func (p *Profile) SocialNetworks() Profile {
	// err := db.GORM().Model(&p).Association("SocialNetwork")
	db.GORM().Preload("SocialNetwork").First(&p)

	return *p
}
