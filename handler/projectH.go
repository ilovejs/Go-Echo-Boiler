package handler

import (
	"github.com/labstack/echo"
	"net/http"
	. "onsite/dto/projects"
	m "onsite/models"
	. "onsite/utils"
)

func (h *Handler) ReadProject(c echo.Context) error {
	panic("no read project")
	//id, err := strconv.Atoi(c.QueryParam("id"))
	//p, err := h.projectStore.Read(id)
	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	//}
	//if p == nil {
	//	return c.JSON(http.StatusNotFound, utils.NotFound())
	//}
	//return c.JSON(http.StatusOK, NewReadProjectResponse(p))
}

func (h *Handler) CreateProject(c echo.Context) error {
	var p m.Project
	req := &CreateProjectRequest{}
	if err := req.Bind(c, &p); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}

	err := h.projectStore.Create(&p)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	return c.JSON(http.StatusCreated, NewProjectResponse(&p))
}
