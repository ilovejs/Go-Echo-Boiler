package users

import (
	m "onsite/models"
	"onsite/utils"
)

type UserResponse struct {
	User struct {
		ID         int    `json:"id"`
		UserRoleID int    `json:"user_role_id"`
		Username   string `json:"username"`
		Email      string `json:"email"`
		Token      string `json:"token,omitempty"`
		IsDeleted  bool   `json:"is_deleted"`
		IsActive   bool   `json:"is_active"`
		Role       *Role  `json:"role"` // object type in json
	} `json:"result"`
}

func NewUserResponse(u *m.User) *UserResponse {
	r := new(UserResponse)
	r.User.ID = u.ID
	r.User.UserRoleID = u.UserRoleID
	r.User.Email = u.Email
	// for each login action, we generate token
	r.User.Token = utils.GenerateJWT(u.ID)
	r.User.Username = u.Username
	r.User.IsDeleted = u.IsDeleted
	r.User.IsActive = u.IsActive

	return r
}

type Role struct {
	ID          string        `json:"id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Describe    string        `json:"describe,omitempty"`
	Permissions []*Permission `json:"permissions,omitempty"`
}

type Permission struct {
}

/* Login */
type UserLoginResponse struct {
	User struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Deleted  int    `json:"deleted"`
		RoleID   string `json:"roleId"`
		Role     *Role  `json:"role"`
		Token    string `json:"token"`
	} `json:"result"`
}

func NewUserLoginResponse(u *m.User) *UserLoginResponse {
	r := new(UserLoginResponse)
	r.User.ID = u.ID
	r.User.Name = "Michael" // TODO: review this code
	r.User.Username = u.Username
	if u.IsDeleted {
		r.User.Deleted = 1
	} else {
		r.User.Deleted = 0
	}
	r.User.RoleID = "admin"
	r.User.Role = &Role{} // HACK: making sure {} not null
	r.User.Token = utils.GenerateJWT(u.ID)
	return r
}

/* List */
type UserListResponse struct {
	Users []*UserResponse `json:"users"`
	Count int             `json:"count"`
}

func NewListUserResponse(users []*m.User, count int) *UserListResponse {
	r := new(UserListResponse)
	for _, u := range users {
		item := NewUserResponse(u)
		item.User.Token = ""
		r.Users = append(r.Users, item)
	}
	r.Count = count
	return r
}

/* Logout */
type UserLogoutResponse struct {
	Message string `json:"message"`
}

func NewUserLogoutResponse() *UserLogoutResponse {
	lr := new(UserLogoutResponse)
	lr.Message = "Logout Success"
	return lr
}

/* Profile */
type ProfileResponse struct {
	Profile struct {
		Username string  `json:"username"`
		Bio      *string `json:"bio"`
		Image    *string `json:"image"`
	} `json:"profile"`
}

func NewProfileResponse(u *m.User) *ProfileResponse {
	r := new(ProfileResponse)
	r.Profile.Username = u.Username
	return r
}
