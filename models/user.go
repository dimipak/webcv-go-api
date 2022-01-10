package models

import (
	dbs "app/repositories/db"
	systemservice "app/systemService"
	"time"
)

type User struct {
	UserId      int    `json:"user_id" gorm:"primarykey"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	ActivateKey string `json:"activate_key"`
	RecoveryKey string `json:"recovery_key"`
	Activated   bool   `json:"activated"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
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
		ActivateKey: systemservice.GetRandomString(32),
		Password:    systemservice.HashAndSalt(password),
	}

	u.SetTimestamps()

	return u
}

func (u *User) Activate() {

	dbs.New().Model(u).Updates(User{
		Activated:   true,
		ActivateKey: "-",
	})
}
