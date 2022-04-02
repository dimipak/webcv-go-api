package models

import (
	db "app/system"
)

type socialNetworkMethods interface {
	Update(socialNetwork SocialNetwork)
}

type SocialNetwork struct {
	SocialNetworkId int    `json:"social_network_id" gorm:"primarykey"`
	ProfileId       int    `json:"profile_id"`
	Linkedin        string `json:"linkedin"`
	Github          string `json:"github"`
	Stackoverflow   string `json:"stackoverflow"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

func (s *SocialNetwork) Update(socialNetwork SocialNetwork) {
	db.GORM().Model(s).Updates(socialNetwork)
}
