package handler

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
	"net/http"
	. "onsite/dto/basic_trade"
	m "onsite/models"
	. "onsite/utils"
	"strconv"
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
	var btid int
	var err error
	if btid, err = strconv.Atoi(c.Param("id")); err != nil {
		return errors.New("param: id parsing error")
	}
	fmt.Println("btid: ", btid)
	bt, err := h.basicTradeStore.Get(btid)
	spew.Dump(bt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if bt == nil {
		return c.JSON(http.StatusNotFound, NewError(err))
	}
	return c.JSON(http.StatusOK, NewBasicTradeResponse(bt[0]))
}

func (h *Handler) ListBasicTrades(c echo.Context) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	tradeItems, count, err := h.basicTradeStore.List(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusOK, NewBasicTradeListResponse(tradeItems, count))
}

func (h *Handler) UpdateBasicTrade(c echo.Context) error {
	panic("no")
}

func (h *Handler) DeleteBasicTrade(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	//del without find
	err = h.basicTradeStore.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusOK, NewDeleteBasicTradeResponse(id))
}
