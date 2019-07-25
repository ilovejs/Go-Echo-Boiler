package handler

import (
	"github.com/labstack/echo"
	"net/http"
	. "onsite/dto/basic_trade"
	m "onsite/models"
	. "onsite/utils"
)

func (h *Handler) CreateBasicTrade(c echo.Context) error {
	var bt m.BasicTrade
	req := &CreateBasicTradeRequest{}
	if err := req.Bind(c, &bt); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	err := h.basicTradeStore.Create(&bt)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	return c.JSON(http.StatusCreated, NewBasicTradeResponse(&bt))
}

func (h *Handler) ReadBasicTrade(c echo.Context) error {
	panic("no")
}

func (h *Handler) UpdateBasicTrade(c echo.Context) error {

	panic("no")
}

func (h *Handler) DeleteBasicTrade(c echo.Context) error {

	panic("no")
}
