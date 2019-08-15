package trades

import (
	"github.com/labstack/echo"
	"onsite/models"
	"time"
)

/* create */
type CreateTradeRequest struct {
	CategoryID int     `json:"category_id" validate:"required"`
	CreatorID  int     `json:"surveyor_id" validate:"required"`
	ProjectID  int     `json:"project_id" validate:"required"`
	Level      string  `json:"level"`
	Subtitle   string  `json:"subtitle" validate:"required"`
	Value      float64 `json:"value"`
}

func (r *CreateTradeRequest) Bind(c echo.Context, t *models.Trade) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	t.TradeCategoryID = r.CategoryID
	t.CreatorID = r.CreatorID
	t.ProjectID = r.ProjectID
	t.Level = r.Level
	t.Subtitle.SetValid(r.Subtitle)
	t.Value.SetValid(r.Value)
	t.Created.SetValid(time.Now())
	return nil
}

/* update */
type UpdateTradeRequest struct {
	CategoryID int     `json:"category_id" validate:"required"`
	CreatorID  int     `json:"creator_id" validate:"required"`
	ProjectID  int     `json:"project_id" validate:"required"`
	Level      string  `json:"level"`
	Subtitle   string  `json:"subtitle" validate:"required"`
	Value      float64 `json:"value"`
	Editable   bool    `json:"editable,omitempty"`
	IsDeleted  bool    `json:"is_deleted,omitempty"`
}

func (r *UpdateTradeRequest) Bind(c echo.Context, t *models.Trade) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	t.TradeCategoryID = r.CategoryID
	t.CreatorID = r.CreatorID
	t.ProjectID = r.ProjectID
	t.Level = r.Level
	t.Subtitle.SetValid(r.Subtitle)
	t.Value.SetValid(r.Value)
	t.Editable = r.Editable
	t.IsDeleted = r.IsDeleted
	t.Updated.SetValid(time.Now())
	return nil
}

/* List request ? */
