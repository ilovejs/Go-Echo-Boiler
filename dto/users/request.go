package users

import (
	"github.com/labstack/echo"
	"github.com/volatiletech/null"
	m "onsite/models"
	. "onsite/utils"
)

type userUpdateRequest struct {
	User struct {
		UserRoleId string `json:"user_role_id" validate:"required"`
		UserName   string `json:"username" validate:"required"`
		Password   string `json:"password" validate:"required"`
		Email      string `json:"email" validate:"email"`
	} `json:"user"`
}

func NewUserUpdateRequest() *userUpdateRequest {
	return new(userUpdateRequest)
}

// take data from model
func (r *userUpdateRequest) Extract(u *m.User) {
	r.User.UserRoleId = string(u.UserRoleID)

	//err := u.Username.Scan(&r.User.UserName)
	//LogIf(err)
	r.User.UserName = u.Username

	err := u.Password.Scan(&r.User.Password)
	LogIf(err)

	u.Email = r.User.Email
}

//custom binder https://echo.labstack.com/guide/request
//mutate model object so as to be saved in next
func (r *userUpdateRequest) Bind(c echo.Context, u *m.User) error {
	// calling default bind
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	//u.Username.SetValid(r.User.UserName)
	u.Username = r.User.UserName

	u.Email = r.User.Email

	// change password
	if password := new(string); u.Password.Valid && u.Password.Scan(&password) != nil {
		// not match previous password
		if r.User.Password != *password {
			h, err := HashPassword(r.User.Password)
			if err != nil {
				return err
			}
			u.Password.SetValid(h)
		}
	}
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
	u.Password = null.StringFrom(hashed)
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
