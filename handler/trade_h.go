package handler

import (
	"github.com/labstack/echo"
	"net/http"
	. "onsite/dto/trades"
	m "onsite/models"
)

func (h *Handler) CreateTrade(c echo.Context) error {
	var t m.Trade
	req := &CreateTradeRequest{}
	if err := req.Bind(c, &t); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	err := h.TradeStore.Create(&t)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	return c.JSON(http.StatusCreated, NewTradeResponse(&t))
}

func (h *Handler) ListTrade(c echo.Context) error {
	panic("no")
}

func (h *Handler) ReadTrade(c echo.Context) error {
	panic("no")
}

func (h *Handler) UpdateTrade(c echo.Context) error {
	panic("no")
}

func (h *Handler) DeleteTrade(c echo.Context) error {
	panic("no")
}
