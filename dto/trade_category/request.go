package trade_category

import (
	"github.com/labstack/echo"
	"onsite/models"
	"time"
)

type CreateTradeCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

func (r *CreateTradeCategoryRequest) Bind(c echo.Context, bt *models.TradeCategory) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	bt.Name = r.Name
	return nil
}

/* update */
type UpdateTradeCategoryRequest struct {
	Name string `json:"name" validate:"required"`
	//todo: we skip to use this field, why keep here ?
	IsActive bool `json:"is_active,omitempty"`
}

func (r *UpdateTradeCategoryRequest) Bind(c echo.Context, item *models.TradeCategory) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	item.Name = r.Name
	item.Updated.SetValid(time.Now())
	return nil
}
