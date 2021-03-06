package repositories

import (
	m "app/models"
	db "app/system"
	"errors"
	"fmt"
)

type UserRepository struct {
	UserId int
}

func (u *UserRepository) GetById() (m.User, error) {
	var user m.User

	err := db.GORM().Where("user_id = ?", u.UserId).First(&user).Error

	return user, err
}

func (u *UserRepository) SetUserId(id int) *UserRepository {
	u.UserId = id
	return u
}

func CreateUser(u m.User) m.User {

	nu := u

	db.GORM().Create(&nu)

	return nu
}

func GetUserByEmail(e string) (m.User, error) {
	var user m.User

	err := db.GORM().First(&user, "email = ?", e).Error

	return user, err
}

func GetUserByUsername(u string) (m.User, error) {
	var user m.User

	err := db.GORM().First(&user, "username = ?", u)
	if err.Error != nil {
		fmt.Println("error!!!")
		return user, errors.New("SQL ERROR")
	}

	return user, nil
}

func GetUserByActivateKey(k string) (m.User, error) {
	var user m.User

	err := db.GORM().First(&user, "activate_key = ?", k)
	if err.Error != nil {
		fmt.Println("error!!!")
		return user, errors.New("SQL ERROR")
	}

	return user, nil
}
