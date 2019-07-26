package handler

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
	"net/http"
	. "onsite/dto/trade_category"
	m "onsite/models"
	. "onsite/utils"
	"strconv"
)

func (h *Handler) CreateTradeCategory(c echo.Context) error {
	var bt m.TradeCategory
	req := &CreateTradeCategoryRequest{}
	if err := req.Bind(c, &bt); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	err := h.TradeCategoryStore.Create(&bt)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	return c.JSON(http.StatusCreated, NewTradeCategoryResponse(&bt))
}

func (h *Handler) ReadTradeCategory(c echo.Context) error {
	var btid int
	var err error
	if btid, err = strconv.Atoi(c.Param("id")); err != nil {
		return errors.New("param: id parsing error")
	}
	fmt.Println("btid: ", btid)
	bt, err := h.TradeCategoryStore.Get(btid)
	spew.Dump(bt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if bt == nil {
		return c.JSON(http.StatusNotFound, NewError(err))
	}
	return c.JSON(http.StatusOK, NewTradeCategoryResponse(bt[0]))
}

func (h *Handler) ListTradeCategories(c echo.Context) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	tradeItems, count, err := h.TradeCategoryStore.List(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusOK, NewTradeCategoryListResponse(tradeItems, count))
}

func (h *Handler) UpdateTradeCategory(c echo.Context) error {
	req := &UpdateTradeCategoryRequest{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.New("parsing id error"))
	}
	arr, err := h.TradeCategoryStore.Get(id)
	if len(arr) == 0 {
		return c.JSON(http.StatusNotFound, NotFound())
	}
	itemToUpdate := arr[0]
	//update model payload
	if err := req.Bind(c, itemToUpdate); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	//commit updating
	err = h.TradeCategoryStore.Update(itemToUpdate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	fmt.Println("TradeCategoryH: Updated")
	return c.JSON(http.StatusOK, NewUpdateTradeCategoryResponse(itemToUpdate))
}

func (h *Handler) DeleteTradeCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, NewError(err))
	}
	//del without find
	err = h.TradeCategoryStore.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	return c.JSON(http.StatusOK, NewDeleteTradeCategoryResponse(id))
}
