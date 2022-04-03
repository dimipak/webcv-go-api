package models

import (
	"app/helpers"
	"app/system"
	db "app/system"

	"time"
)

type User struct {
	UserId      int      `json:"user_id" gorm:"primarykey"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Email       string   `json:"email"`
	ActivateKey string   `json:"activate_key"`
	RecoveryKey string   `json:"recovery_key"`
	Activated   bool     `json:"activated"`
	Profile     Profiles `json:"profiles" gorm:"foreignKey:UserId;references:UserId`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

const (
	timeFormat string = "2006-01-02 15:04:05"
)

func (u *User) SetCreatedAt() {
	u.CreatedAt = time.Now().Format(timeFormat)
}

func (u *User) SetUpdatedAt() {
	u.UpdatedAt = time.Now().Format(timeFormat)
}

func (u *User) SetTimestamps() {
	u.SetCreatedAt()
	u.SetUpdatedAt()
}

func NewUser(username string, email string, password string) User {

	u := User{
		Username:    username,
		Email:       email,
		ActivateKey: helpers.GetRandomString(32),
		Password:    system.HashAndSalt(password),
	}

	u.SetTimestamps()

	return u
}

func (u *User) Activate() {

	db.GORM().Model(u).Updates(User{
		Activated:   true,
		ActivateKey: "-",
	})
}

func (u *User) Profiles() User {
	db.GORM().Preload("Profile").First(&u)
	return *u
}
