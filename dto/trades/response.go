package trades

import (
	"github.com/volatiletech/null"
	"onsite/models"
)

type TradeResponse struct {
	KEY        int       `json:"key"`
	ID         int       `json:"id"`
	ProjectID  int       `json:"pid"`
	CreatorID  int       `json:"cid"`
	Category   string    `json:"cat,omitempty"`
	CategoryID int       `json:"cat_id"`
	Subtitle   string    `json:"subtitle"`
	Editable   bool      `json:"editable"`
	Created    null.Time `json:"created"`
}

/* Create */
func NewTradeResponse(t *models.Trade, category string) *TradeResponse {
	resp := new(TradeResponse)
	resp.KEY = t.ID
	resp.ID = t.ID
	resp.ProjectID = t.ProjectID
	resp.CreatorID = t.CreatorID

	resp.Category = category
	resp.CategoryID = t.TradeCategoryID
	resp.Subtitle = *t.Subtitle.Ptr()
	resp.Created = t.Created
	return resp
}

type TradeListResponse struct {
	PageSize   int              `json:"pageSize"`
	PageNo     int              `json:"pageNo"`
	TotalCount int              `json:"totalCount"`
	TotalPage  int              `json:"totalPage"`
	Trades     []*TradeResponse `json:"data"`
}

/* List */
func NewTradeListResponse(
	trades []*models.Trade,
	totalPage int,
	totalCnt int) *TradeListResponse {

	r := new(TradeListResponse)
	for _, t := range trades {
		// reuse
		item := NewTradeResponse(t, t.R.TradeCategory.Name)
		r.Trades = append(r.Trades, item)
	}
	r.TotalCount = totalCnt
	r.TotalPage = totalPage
	return r
}

/* Update */
type UpdateTradeResponse struct {
	ID         int     `json:"id"`
	CategoryID int     `json:"category_id"`
	CreatorID  int     `json:"creator_id"`
	ProjectID  int     `json:"project_id"`
	Level      string  `json:"level"`
	Subtitle   string  `json:"subtitle"`
	Value      float64 `json:"value"`
	Editable   bool    `json:"editable"`
	IsDeleted  bool    `json:"is_deleted"`
}

// fn
func NewUpdateTradeResponse(t *models.Trade) *UpdateTradeResponse {
	resp := new(UpdateTradeResponse)
	resp.ID = t.ID
	resp.CategoryID = t.TradeCategoryID
	resp.CreatorID = t.CreatorID
	resp.ProjectID = t.ProjectID
	resp.Level = t.Level
	resp.Subtitle = *t.Subtitle.Ptr()
	resp.Value = *t.Value.Ptr()
	resp.Editable = t.Editable
	resp.IsDeleted = t.IsDeleted
	return resp
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
