package trade_categories

import (
	"github.com/volatiletech/null"
	"onsite/models"
	"time"
)

// Create
type TradeCategoryResponse struct {
	ID      int       `json:"trade_category_id"`
	Name    string    `json:"name"`
	Created null.Time `json:"created"`
}

func NewTradeCategoryResponse(bt *models.TradeCategory) *TradeCategoryResponse {
	resp := new(TradeCategoryResponse)
	resp.ID = bt.ID
	resp.Name = bt.Name
	resp.Created = bt.Created
	return resp
}

/* Update */
type UpdateTradeCategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewUpdateTradeCategoryResponse(item *models.TradeCategory) *UpdateTradeCategoryResponse {
	resp := new(UpdateTradeCategoryResponse)
	resp.Name = item.Name
	return resp
}

/* List */
type TradeCategoryListResponse struct {
	TradeCategories []*TradeCategoryResponse `json:"trade_categories"`
	ItemCount       int                      `json:"item_count"`
}

func NewTradeCategoryListResponse(
	TradeCategoriesSource []*models.TradeCategory,
	count int) *TradeCategoryListResponse {

	r := new(TradeCategoryListResponse)
	r.TradeCategories = make([]*TradeCategoryResponse, 0)
	timeNow := time.Now()

	for _, btSrc := range TradeCategoriesSource {
		item := new(TradeCategoryResponse)
		item.ID = btSrc.ID
		item.Name = btSrc.Name
		item.Created.SetValid(timeNow)

		r.TradeCategories = append(r.TradeCategories, item)
	}
	r.ItemCount = count
	return r
}

/* Delete */
type DeleteTradeCategoryResponse struct {
	TradeCategory struct {
		ID int `json:"id"`
	} `json:"trade_category"`
}

func NewDeleteTradeCategoryResponse(id int) *DeleteTradeCategoryResponse {
	resp := new(DeleteTradeCategoryResponse)
	resp.TradeCategory.ID = id
	return resp
}
