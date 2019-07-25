package basic_trade

import (
	"github.com/volatiletech/null"
	"onsite/models"
)

type BasicTradeResponse struct {
	ID      int       `json:"basic_trade_id"`
	Created null.Time `json:"created"`
}

func NewBasicTradeResponse(bt *models.BasicTrade) *BasicTradeResponse {
	resp := new(BasicTradeResponse)
	resp.ID = bt.ID
	resp.Created = bt.Created
	return resp
}
