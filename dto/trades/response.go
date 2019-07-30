package trades

import (
	"github.com/volatiletech/null"
	"onsite/models"
)

type TradeResponse struct {
	ID      int       `json:"trade_id"`
	Desc    string    `json:"description"`
	Created null.Time `json:"created"`
}

type TradeListResponse struct {
	Trades []*TradeResponse `json:"trades"`
	Count  int              `json:"count"`
}

/* create */
func NewTradeResponse(t *models.Trade) *TradeResponse {
	resp := new(TradeResponse)
	resp.ID = t.ID
	resp.Desc = *t.Description.Ptr()
	resp.Created = t.Created
	return resp
}

/* Update */
type UpdateTradeResponse struct {
	ID         int     `json:"id"`
	CategoryID int     `json:"category_id"`
	SurveyorID int     `json:"surveyor_id"`
	ProjectID  int     `json:"project_id"`
	Level      string  `json:"level"`
	Desc       string  `json:"description"`
	Value      float64 `json:"value"`
	IsActive   bool    `json:"is_active"`
	IsDeleted  bool    `json:"is_deleted"`
}

func NewUpdateTradeResponse(t *models.Trade) *UpdateTradeResponse {
	resp := new(UpdateTradeResponse)
	resp.ID = t.ID
	resp.CategoryID = t.TradeCategoryID
	resp.SurveyorID = t.SurveyorID
	resp.ProjectID = t.ProjectID
	resp.Level = t.Level
	resp.Desc = *t.Description.Ptr()
	resp.Value = *t.Value.Ptr()
	resp.IsActive = t.IsActive
	resp.IsDeleted = t.IsDeleted
	return resp
}

/* List */
func NewTradeListResponse(trades []*models.Trade, count int) *TradeListResponse {
	r := new(TradeListResponse)
	for _, t := range trades {
		// reuse
		item := NewTradeResponse(t)
		r.Trades = append(r.Trades, item)
	}
	r.Count = count
	return r
}

/* Delete */
type DeleteTradeResponse struct {
	Trade struct {
		ID int `json:"id"`
	} `json:"trade"`
}

func NewDeleteTradeResponse(id int) *DeleteTradeResponse {
	resp := new(DeleteTradeResponse)
	resp.Trade.ID = id
	return resp
}
