package resources

import (
	"app/models"
	"time"
)

type User struct {
	UserId    int    `json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func UserResource(u models.User) User {

	layout := "2006-01-02T15:04:05Z"
	format := "2006-01-02 15:04:05"

	createdAt, _ := time.Parse(layout, u.CreatedAt)
	updatedAt, _ := time.Parse(layout, u.UpdatedAt)

	return User{
		UserId:    u.UserId,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: createdAt.Format(format),
		UpdatedAt: updatedAt.Format(format),
	}
}
