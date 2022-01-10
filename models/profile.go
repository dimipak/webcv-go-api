package models

import dbs "app/repositories/db"

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

func (p *Profile) Activate() {

	dbs.New().Model(p).Updates(Profile{
		Active: true,
	})
}

func (p *Profile) DeActivate() {

	dbs.New().Model(p).Updates(map[string]interface{}{
		"active": false,
	})
}
