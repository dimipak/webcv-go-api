package helpers

import (
	"math/rand"
	"time"
)

const pool string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func GetRandomString(n int) string {
	var letters = []rune(pool)

	s := make([]rune, n)

	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}
