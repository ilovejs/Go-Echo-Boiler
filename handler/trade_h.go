package handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"math"
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
	return c.JSON(http.StatusCreated, NewTradeResponse(&t, ""))
}

func (h *Handler) ReadTrade(c echo.Context) error {
	var tid int
	var err error
	if tid, err = strconv.Atoi(c.Param("id")); err != nil {
		return errors.New("param: id parsing error")
	}
	fmt.Println("tid: ", tid)

	// We can supply n ids, Get can handle that
	t, err := h.TradeStore.Get(tid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if t == nil {
		return c.JSON(http.StatusNotFound, NewError(err))
	}
	return c.JSON(http.StatusOK, NewTradeResponse(t[0], t[0].R.TradeCategory.Name))
}

func (h *Handler) ListTrade(c echo.Context) error {
	// NOTE: cast to int so make it little bit safe
	projectId, err := strconv.Atoi(c.QueryParam("pid"))
	if err != nil {
		// todo: server log
		fmt.Println("pid parse error", err)
	}
	fmt.Println("Pid ========>", projectId)

	pageNo, err := strconv.Atoi(c.QueryParam("pageNo"))
	if err != nil {
		pageNo = 1
		fmt.Println("pageNo parse error", err)
	}
	fmt.Println("pageNo", pageNo)

	pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
	if err != nil {
		// default page size
		pageSize = 10
		fmt.Println("pageSize parse error", err)
	}
	fmt.Println("pageSize", pageSize)

	offset := (pageNo - 1) * pageSize

	tradeItems, totalCount, err := h.TradeStore.List(
		projectId,
		offset,
		pageSize) // offset, limit
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	totalPage := int(math.Ceil(
		float64(totalCount) / float64(pageSize),
	))
	return c.JSON(http.StatusOK, NewTradeListResponse(
		tradeItems,
		pageSize, pageNo,
		totalPage, totalCount))
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
