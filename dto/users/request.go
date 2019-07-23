package users

import (
	"github.com/labstack/echo"
	m "onsite/models"
	. "onsite/utils"
)

type userUpdateRequest struct {
	User struct {
		UserRoleId int    `json:"user_role_id" validate:"required"`
		UserName   string `json:"username" validate:"required"`
		Password   string `json:"password" validate:"required"`
		Email      string `json:"email" validate:"email"`
		IsDeleted  bool   `json:"is_deleted"`
		IsActive   bool   `json:"is_active"`
	} `json:"user"`
}

func NewUserUpdateRequest() *userUpdateRequest {
	return new(userUpdateRequest)
}

// take data from model
func (r *userUpdateRequest) LoadFrom(u *m.User) {
	r.User.UserRoleId = u.UserRoleID
	r.User.UserName = u.Username
	r.User.Password = u.Password
	r.User.Email = u.Email
	r.User.IsDeleted = u.IsDeleted
	r.User.IsActive = u.IsActive
}

// Custom binder https://echo.labstack.com/guide/request
// Mutate model object so as to be saved in next
func (r *userUpdateRequest) Bind(c echo.Context, u *m.User) error {
	var err error
	// calling default bind
	if err = c.Bind(r); err != nil {
		return err
	}
	if err = c.Validate(r); err != nil {
		return err
	}
	// fields to update in model
	u.UserRoleID = r.User.UserRoleId
	u.Username = r.User.UserName

	// change password due to not match previous password
	if r.User.Password != u.Password {
		h, err := HashPassword(r.User.Password)
		if err != nil {
			return err
		}
		u.Password = h
	}

	// todo: Don't update email, cuz we don't send email to let user reset password themselves yet
	//  email is unique so safe from impersonating a different user
	//u.Email = r.User.Email

	u.IsDeleted = r.User.IsDeleted
	u.IsActive = r.User.IsActive
	return nil
}

/* Register */
type UserRegisterRequest struct {
	User struct {
		UserRoleId int    `json:"user_role_id" validate:"required"`
		Username   string `json:"username" validate:"required"`
		Email      string `json:"email" validate:"required,email"`
		Password   string `json:"password" validate:"required"`
	} `json:"user"`
}

//only touch 2 fields: 'new' username and email
func (r *UserRegisterRequest) Bind(c echo.Context, u *m.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.UserRoleID = r.User.UserRoleId
	//u.Username = null.StringFrom(r.User.Username)
	u.Username = r.User.Username
	u.Email = r.User.Email
	hashed, err := HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Password = hashed
	return nil
}

/* Login */
type UserLoginRequest struct {
	User struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

// only validate, not to touch model
func (r *UserLoginRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}
