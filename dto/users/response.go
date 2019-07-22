package users

import (
	. "onsite/models"
	"onsite/utils"
)

type UserResponse struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	} `json:"user"`
}

func NewUserResponse(u *User) *UserResponse {
	r := new(UserResponse)

	err := u.Username.Scan(r.User.Username)
	utils.LogIf(err)

	r.User.Email = u.Email
	r.User.Token = utils.GenerateJWT(u.ID)
	return r
}

type ProfileResponse struct {
	Profile struct {
		Username string  `json:"username"`
		Bio      *string `json:"bio"`
		Image    *string `json:"image"`
	} `json:"profile"`
}

func NewProfileResponse(u User) *ProfileResponse {
	r := new(ProfileResponse)
	err := u.Username.Scan(r.Profile.Username)
	utils.LogIf(err)
	return r
}
