package handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	. "onsite/dto/trades"
	m "onsite/models"
	. "onsite/utils"
	"strconv"
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

func (h *Handler) ReadTrade(c echo.Context) error {
	var tid int
	var err error
	if tid, err = strconv.Atoi(c.Param("id")); err != nil {
		return errors.New("param: id parsing error")
	}
	fmt.Println("tid: ", tid)

	t, err := h.TradeStore.Get(tid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if t == nil {
		return c.JSON(http.StatusNotFound, NewError(err))
	}
	return c.JSON(http.StatusOK, NewTradeResponse(t[0]))
}

func (h *Handler) ListTrade(c echo.Context) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	tradeItems, count, err := h.TradeStore.List(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusOK, NewTradeListResponse(tradeItems, count))
}

func (h *Handler) UpdateTrade(c echo.Context) error {
	req := &UpdateTradeRequest{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.New("parsing id error"))
	}
	arr, err := h.TradeStore.Get(id)
	if len(arr) == 0 {
		return c.JSON(http.StatusNotFound, NotFound())
	}
	itemToUpdate := arr[0]
	// update model payload
	if err := req.Bind(c, itemToUpdate); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	// commit updating
	err = h.TradeStore.Update(itemToUpdate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	fmt.Println("TradeH: Updated")
	return c.JSON(http.StatusOK, NewUpdateTradeResponse(itemToUpdate))
}

func (h *Handler) DeleteTrade(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	// del without find
	err = h.TradeStore.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusOK, NewDeleteTradeResponse(id))
}
