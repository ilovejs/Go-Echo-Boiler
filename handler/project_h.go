package handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	. "onsite/dto/projects"
	m "onsite/models"
	. "onsite/utils"
	"strconv"
)

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

func (h *Handler) ReadProject(c echo.Context) error {
	// get this from url.go where router is defined
	pid, err := strconv.Atoi(c.Param("id"))
	fmt.Println("Pid is: ", pid)
	p, err := h.projectStore.Read(pid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if p == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}
	return c.JSON(http.StatusOK, NewProjectResponse(p))
}

func (h *Handler) ListProject(c echo.Context) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	projects, count, err := h.projectStore.List(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusOK, NewProjectListResponse(projects, count))
}

func (h *Handler) UpdateProject(c echo.Context) error {
	req := &UpdateProjectRequest{}

	pid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.New("id wrong format or not given"))
	}
	p, err := h.projectStore.Read(pid)
	if p == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}
	// update model payload
	if err := req.Bind(c, p); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	// commit updating
	err = h.projectStore.Update(p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	fmt.Println("ProjectH: Updated")
	return c.JSON(http.StatusOK, NewUpdateProjectResponse(p))
}

func (h *Handler) DeleteProject(c echo.Context) error {
	// not QueryParam
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	// find
	p, err := h.projectStore.Read(id)
	if p == nil {
		return c.JSON(http.StatusNotFound, NotFound())
	}
	// del
	err = h.projectStore.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusOK, NewDeleteProjectResponse(p))
}
