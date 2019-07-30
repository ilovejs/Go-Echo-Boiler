package trades

import (
	"github.com/labstack/echo"
	"onsite/models"
	"time"
)

/* create */
type CreateTradeRequest struct {
	Desc       string  `json:"description" validate:"required"`
	ProjectID  int     `json:"project_id" validate:"required"`
	CategoryID int     `json:"category_id" validate:"required"`
	SurveyorID int     `json:"surveyor_id" validate:"required"`
	Level      string  `json:"level"`
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
	t.SurveyorID = r.SurveyorID
	t.ProjectID = r.ProjectID
	t.Level = r.Level
	t.Description.SetValid(r.Desc)
	t.Value.SetValid(r.Value)
	t.Created.SetValid(time.Now())
	return nil
}

/* update */
type UpdateTradeRequest struct {
	CategoryID int     `json:"category_id" validate:"required"`
	SurveyorID int     `json:"surveyor_id" validate:"required"`
	ProjectID  int     `json:"project_id" validate:"required"`
	Level      string  `json:"level"`
	Desc       string  `json:"description" validate:"required"`
	Value      float64 `json:"value"`
	IsActive   bool    `json:"is_active,omitempty"`
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
	t.SurveyorID = r.SurveyorID
	t.ProjectID = r.ProjectID
	t.Level = r.Level
	t.Description.SetValid(r.Desc)
	t.Value.SetValid(r.Value)
	t.IsActive = r.IsActive
	t.IsDeleted = r.IsDeleted
	t.Updated.SetValid(time.Now())
	return nil
}
