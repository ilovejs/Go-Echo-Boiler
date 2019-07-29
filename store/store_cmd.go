package store

import (
	"onsite/models"
)

/* Store interfaces are injected to handler in url.go */

type UserStoreInterface interface {
	Create(*models.User) error
	GetByID(int) (*models.User, error)
	GetByEmail(string) (*models.User, error)
	GetByUserName(string) (*models.User, error)
	Update(*models.User) error
	Delete(int) error
}

type ProjectStoreInterface interface {
	Create(*models.Project) error
	Read(int) (*models.Project, error)
	Update(*models.Project) error
	Delete(int) error
	List(offset int, limit int) ([]*models.Project, int, error)
	Contractors(pid int) ([]*models.User, error)
}

type TradeCategoryStoreInterface interface {
	Create(trade *models.TradeCategory) error
	Get(ids ...int) ([]*models.TradeCategory, error) //Get one or all
	Update(trade *models.TradeCategory) error
	Delete(int) error

	List(offset int, limit int) ([]*models.TradeCategory, int, error)
}

type TradeStoreInterface interface {
	Create(trade *models.Trade) error
	Get(ids ...int) ([]*models.Trade, error) //Get one or all
	Update(trade *models.Trade) error
	Delete(int) error
	List(offset int, limit int) ([]*models.Trade, int, error)
	GetByProject(projectId int) ([]*models.Trade, error)
	Count(tradeId) (int, err)
}

type ClaimStoreInterface interface {
	//
}
