package dto

import (
	"github.com/labstack/echo/v4"
	"onsite/model"
)

type userUpdateRequest struct {
	User struct {
		UserName 	string `json:"username"`
		FirstName   string `json:"firstname"`
		LastName    string `json:"lastname"`
		Company     string `json:"company"`
		Email    	string `json:"email" validate:"email"`
		Mobile      string `json:"mobile"`
		Password 	string `json:"password"`
		Bio      	string `json:"bio"`
		Image    	string `json:"image"`
	} `json:"user"`
}

func NewUserUpdateRequest() *userUpdateRequest {
	return new(userUpdateRequest)
}

//populate request by model
func (r *userUpdateRequest) Populate(u *model.User) {
	r.User.UserName = u.UserName
	r.User.FirstName = u.FirstName
	r.User.LastName = u.LastName
	r.User.Company = u.Company
	r.User.Email = u.Email
	r.User.Mobile = u.Mobile
	r.User.Password = u.Password
	if u.Bio != nil {
		r.User.Bio = *u.Bio
	}
	if u.Image != nil {
		r.User.Image = *u.Image
	}
}

//custom binder https://echo.labstack.com/guide/request
//mutate model object so as to be saved in next
func (r *userUpdateRequest) Bind(c echo.Context, u *model.User) error {
	//still calling default bind
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.UserName = r.User.UserName
	u.FirstName = r.User.FirstName
	u.LastName = r.User.LastName
	u.Company = r.User.Company
	u.Email = r.User.Email
	u.Mobile = r.User.Mobile
	if r.User.Password != u.Password {
		h, err := u.HashPassword(r.User.Password)
		if err != nil {
			return err
		}
		u.Password = h
	}
	u.Bio = &r.User.Bio
	u.Image = &r.User.Image
	return nil
}

// Register
type UserRegisterRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

//only touch 2 fields: 'new' username and email
func (r *UserRegisterRequest) Bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.UserName = r.User.Username
	u.Email = r.User.Email
	h, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}

// Login
type UserLoginRequest struct {
	User struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

//only validate, not to touch model
func (r *UserLoginRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

