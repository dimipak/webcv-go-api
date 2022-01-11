package models

import (
	dbs "app/repositories/db"
	"time"
)

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
	ProfileImage  string        `json:"profile_image"`
	CoverImage    string        `json:"cover_image"`
	SocialNetwork SocialNetwork `gorm:"foreignKey:ProfileId;references:ProfileId"`
	Skills        []Skill       `gorm:"foreignKey:ProfileId;references:ProfileId"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
}

func nowFormatted() string {
	return time.Now().Format(timeFormat)
}

func (p *Profile) Activate() {

	dbs.New().Model(p).Updates(Profile{
		Active:    true,
		UpdatedAt: nowFormatted(),
	})
}

func (p *Profile) DeActivate() {

	dbs.New().Model(p).Updates(map[string]interface{}{
		"active":     false,
		"updated_at": nowFormatted(),
	})
}

func NewProfile(userId int, username string, firstName string, lastName string) Profile {

	return Profile{
		UserId:    userId,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: nowFormatted(),
		UpdatedAt: nowFormatted(),
	}
}
