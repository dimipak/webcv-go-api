package authentication

import (
	"app/config"
	"app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type mytoken struct {
	Iss    string `json:"iss"`
	UserId int    `json:"user_id"`
	Iat    int    `json:"iat"`
	Exp    int    `json:"exp"`
}

var Auth mytoken

const ISSUER string = "dimipak.gr"

func Sign(u models.User) (string, error) {

	var mySigningKey []byte = []byte(config.Database().JWTSecret)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["iss"] = ISSUER
	claims["user_id"] = u.UserId
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	return token.SignedString(mySigningKey)
}

func Verify(r *http.Request) bool {

	authorizationHeader := r.Header.Get("Authorization")

	if authorizationHeader == "" {
		return false
	}

	accessToken := strings.Split(authorizationHeader, " ")[1]

	var mySigningKey []byte = []byte(config.Database().JWTSecret)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println("parse error = ", err.Error())
		return false
	}

	decode(token)

	return token.Valid
}

func decode(token *jwt.Token) {

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("token failed decoded")
	}

	result, _ := json.Marshal(claims)

	json.Unmarshal(result, &Auth)
}
