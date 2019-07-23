package users

import (
	. "onsite/models"
	"onsite/utils"
	"strconv"
)

type UserResponse struct {
	User struct {
		ID         string `json:"id"`
		UserRoleID string `json:"user_role_id"`
		Username   string `json:"username"`
		Email      string `json:"email"`
		Token      string `json:"token"`
	} `json:"user"`
}

func NewUserResponse(u *User) *UserResponse {
	r := new(UserResponse)
	r.User.ID = strconv.Itoa(u.ID)
	r.User.UserRoleID = strconv.Itoa(u.UserRoleID)
	r.User.Email = u.Email
	//for each login action, we generate token
	r.User.Token = utils.GenerateJWT(u.ID)
	r.User.Username = u.Username
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
	//err := u.Username.Scan(&r.Profile.Username)
	r.Profile.Username = u.Username
	return r
}
