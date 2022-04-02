package system

import (
	"app/config"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd string) string {

	secret := config.G_DATABASE.JWTSecret

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd+secret), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	} // GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func PasswordVerify(password string, hashed string) bool {

	secret := config.G_DATABASE.JWTSecret

	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashed)
	bytePlain := []byte(password + secret)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
