package repositories

import (
	m "app/models"
	db "app/system"
	"errors"
	"fmt"
)

func CreateUser(u m.User) m.User {

	nu := u

	db.GORM().Create(&nu)

	return nu
}

func GetUserByEmail(e string) (m.User, error) {
	var user m.User

	err := db.GORM().First(&user, "email = ?", e)
	if err.Error != nil {
		fmt.Println("error!!!")
		return user, errors.New("SQL ERROR")
	}

	return user, nil
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
