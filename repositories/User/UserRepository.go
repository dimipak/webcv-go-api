package user

import (
	m "app/models"
	r "app/repositories"
	"errors"
	"fmt"
)

func Create(u m.User) m.User {

	nu := u

	r.DB().Create(&nu)

	return nu
}

func GetUserByEmail(e string) (m.User, error) {
	var user m.User

	err := r.DB().First(&user, "email = ?", e)
	if err.Error != nil {
		fmt.Println("error!!!")
		return user, errors.New("SQL ERROR")
	}

	return user, nil
}

func GetUserByUsername(u string) (m.User, error) {
	var user m.User

	err := r.DB().First(&user, "username = ?", u)
	if err.Error != nil {
		fmt.Println("error!!!")
		return user, errors.New("SQL ERROR")
	}

	return user, nil
}

func GetUserByActivateKey(k string) (m.User, error) {
	var user m.User

	err := r.DB().First(&user, "activate_key = ?", k)
	if err.Error != nil {
		fmt.Println("error!!!")
		return user, errors.New("SQL ERROR")
	}

	return user, nil
}
