package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	. "onsite/dto"
	"onsite/model"
	"onsite/utils"
	"strconv"
)

func (h *Handler) ReadProject(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	p, err := h.projectStore.ReadProject(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if p == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return c.JSON(http.StatusOK, NewReadProjectResponse(p))
}

func (h *Handler) CreateProject(c echo.Context) error {
	var p model.Project
	req := &CreateProjectRequest{}
	if err := req.Bind(c, &p); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	created, err := h.projectStore.CreateProject(&p)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, NewCreateProjectResponse(created))
}
