package basic_trade

import (
	"github.com/volatiletech/null"
	"onsite/models"
	"time"
)

//single response
type BasicTradeResponse struct {
	ID      int       `json:"basic_trade_id"`
	Name    string    `json:"name"`
	Created null.Time `json:"created"`
}

func NewBasicTradeResponse(bt *models.BasicTrade) *BasicTradeResponse {
	resp := new(BasicTradeResponse)
	resp.ID = bt.ID
	resp.Name = bt.Name
	resp.Created = bt.Created
	return resp
}

type BasicTradeListResponse struct {
	BasicTrades []*BasicTradeResponse `json:"basic_trades"`
	ItemCount   int                   `json:"item_count"`
}

//list response
func NewBasicTradeListResponse(
	basicTradesSource []*models.BasicTrade,
	count int) *BasicTradeListResponse {

	r := new(BasicTradeListResponse)
	r.BasicTrades = make([]*BasicTradeResponse, 0)
	timeNow := time.Now()

	for _, btSrc := range basicTradesSource {
		item := new(BasicTradeResponse)
		item.ID = btSrc.ID
		item.Name = btSrc.Name
		item.Created.SetValid(timeNow)

		r.BasicTrades = append(r.BasicTrades, item)
	}
	r.ItemCount = count
	return r
}

/* delete */
type DeleteBasicTradeResponse struct {
	BasicTrade struct {
		ID int `json:"id"`
	} `json:"basic_trade"`
}

func NewDeleteBasicTradeResponse(id int) *DeleteBasicTradeResponse {
	resp := new(DeleteBasicTradeResponse)
	resp.BasicTrade.ID = id
	return resp
}
