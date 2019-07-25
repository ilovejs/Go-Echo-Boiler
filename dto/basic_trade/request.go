package basic_trade

import (
	"github.com/labstack/echo"
	"onsite/models"
	"time"
)

type CreateBasicTradeRequest struct {
	Name string `json:"name" validate:"required"`
}

func (r *CreateBasicTradeRequest) Bind(c echo.Context, bt *models.BasicTrade) error {
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
type UpdateBasicTradeRequest struct {
	Name string `json:"name" validate:"required"`
	//todo: we skip to use this field, why keep here ?
	IsActive bool `json:"is_active,omitempty"`
}

func (r *UpdateBasicTradeRequest) Bind(c echo.Context, item *models.BasicTrade) error {
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
