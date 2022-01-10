package models

type SocialNetwork struct {
	SocialNewtorkId int    `json:"social_network_id"`
	ProfileId       int    `json:"profile_id"`
	Linkedin        string `json:"linkedin"`
	Github          string `json:"github"`
	StackOverflow   string `json:"stackoverflow"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
