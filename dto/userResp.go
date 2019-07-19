package dto

import (
	"onsite/model"
	"onsite/utils"
)

type UserResponse struct {
	User struct {
		Username string  `json:"username"`
		Email    string  `json:"email"`
		Bio      *string `json:"bio"`
		Image    *string `json:"image"`
		Token    string  `json:"token"`
	} `json:"user"`
}

func NewUserResponse(u *model.User) *UserResponse {
	r := new(UserResponse)
	r.User.Username = u.UserName
	r.User.Email = u.Email
	r.User.Bio = u.Bio
	r.User.Image = u.Image
	r.User.Token = utils.GenerateJWT(u.ID)
	return r
}

type ProfileResponse struct {
	Profile struct {
		Username  string  `json:"username"`
		Bio       *string `json:"bio"`
		Image     *string `json:"image"`
	} `json:"profile"`
}

func NewProfileResponse(u *model.User) *ProfileResponse {
	r := new(ProfileResponse)
	r.Profile.Username = u.UserName
	r.Profile.Bio = u.Bio
	r.Profile.Image = u.Image
	return r
}

