package users

import (
	. "onsite/models"
	"onsite/utils"
)

type UserResponse struct {
	User struct {
		ID         int    `json:"id"`
		UserRoleID int    `json:"user_role_id"`
		Username   string `json:"username"`
		Email      string `json:"email"`
		Token      string `json:"token"`
		IsDeleted  bool   `json:"is_deleted"`
		IsActive   bool   `json:"is_active"`
	} `json:"user"`
}

func NewUserResponse(u *User) *UserResponse {
	r := new(UserResponse)
	r.User.ID = u.ID
	r.User.UserRoleID = u.UserRoleID
	r.User.Email = u.Email
	//for each login action, we generate token
	r.User.Token = utils.GenerateJWT(u.ID)
	r.User.Username = u.Username
	r.User.IsDeleted = u.IsDeleted
	r.User.IsActive = u.IsActive
	return r
}

type ProfileResponse struct {
	Profile struct {
		Username string  `json:"username"`
		Bio      *string `json:"bio"`
		Image    *string `json:"image"`
	} `json:"profile"`
}

func NewProfileResponse(u *User) *ProfileResponse {
	r := new(ProfileResponse)
	//err := u.Username.Scan(&r.Profile.Username)
	r.Profile.Username = u.Username
	return r
}
