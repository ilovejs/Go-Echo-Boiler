package handler

import (
	"github.com/labstack/echo"
	"net/http"
	. "onsite/dto/users"
	"onsite/models"
	. "onsite/utils"
)

func (h *Handler) SignUp(c echo.Context) error {
	var u models.User
	req := &UserRegisterRequest{}

	if err := req.Bind(c, u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	// create user
	if err := h.userStore.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	return c.JSON(http.StatusCreated, NewUserResponse(&u))
}

func (h *Handler) Login(c echo.Context) error {
	req := &UserLoginRequest{}
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	u, err := h.userStore.GetByEmail(req.User.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, AccessForbidden())
	}
	if !CheckPassword(u, req.User.Password) {
		return c.JSON(http.StatusForbidden, AccessForbidden())
	}
	return c.JSON(http.StatusOK, NewUserResponse(u))
}

func (h *Handler) CurrentUser(c echo.Context) error {
	u, err := h.userStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}
	return c.JSON(http.StatusOK, NewUserResponse(u))
}

func (h *Handler) UpdateUser(c echo.Context) error {
	userToUpdate, err := h.userStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if userToUpdate == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}

	req := NewUserUpdateRequest()
	req.Extract(*userToUpdate)
	//bind context to model struct
	if err := req.Bind(c, *userToUpdate); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	//update
	if err := h.userStore.Update(userToUpdate); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	return c.JSON(http.StatusOK, NewUserResponse(userToUpdate))
}

func (h *Handler) GetProfile(c echo.Context) error {
	username := c.Param("username")
	u, err := h.userStore.GetByUserName(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}
	return c.JSON(http.StatusOK, NewProfileResponse(*u))
}

// utils of handler

func userIDFromToken(c echo.Context) uint {
	id, ok := c.Get("user").(uint)
	if !ok {
		return 0
	}
	return id
}
