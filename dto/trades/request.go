package trades

import (
	"github.com/labstack/echo"
	"onsite/models"
	"time"
)

/* create */
type CreateTradeRequest struct {
	Desc string `json:"name" validate:"required"`
}

func (r *CreateTradeRequest) Bind(
	c echo.Context,
	t *models.Trade) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	t.Description.SetValid(r.Desc)
	return nil
}

/* update */
type UpdateTradeRequest struct {
	Desc string `json:"description" validate:"required"`
	// todo: we skip to use this field, why keep here ?
	IsActive bool `json:"is_active,omitempty"`
}

func (r *UpdateTradeRequest) Bind(
	c echo.Context,
	item *models.Trade) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	item.Description.SetValid(r.Desc)
	item.Updated.SetValid(time.Now())
	return nil
}
