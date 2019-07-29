package trades

import (
	"github.com/volatiletech/null"
	"onsite/models"
	"time"
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

/* update */
type UpdateTradeResponse struct {
	ID   int    `json:"id"`
	Desc string `json:"name"`
}

func NewUpdateTradeResponse(item *models.Trade) *UpdateTradeRequest {
	resp := new(UpdateTradeRequest)
	resp.Desc = *item.Description.Ptr()
	return resp
}

/* List */
func NewTradeListResponse(
	TradesSource []*models.Trade,
	count int) *TradeListResponse {

	r := new(TradeListResponse)
	r.Trades = make([]*TradeResponse, 0)
	timeNow := time.Now()

	for _, btSrc := range TradesSource {
		item := new(TradeResponse)
		item.ID = btSrc.ID
		item.Desc = *btSrc.Description.Ptr()
		item.Created.SetValid(timeNow)

		r.Trades = append(r.Trades, item)
	}
	r.Count = count
	return r
}

/* delete */
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
