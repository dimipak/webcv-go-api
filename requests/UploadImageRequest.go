package requests

import "net/http"

type UploadProfileImageRequest struct {
	ProfileImage string `form:"profile_image" json:"profile_image" validate:"required"`
}

func (u *UploadProfileImageRequest) ValidateRequest(r *http.Request) error {
	return validateRequest(r, u)
}
