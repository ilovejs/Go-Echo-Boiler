package basic_trade

import (
	"github.com/labstack/echo"
	"onsite/models"
)

type CreateBasicTradeRequest struct {
	Name string `json:"name" validate:"required"`

	//IsActive  bool
	//IsDeleted bool
	//Created   null.Time
	//Updated   null.Time
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
