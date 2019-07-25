package store

import (
	"database/sql"
	"onsite/models"
)

type BasicTradeStore struct {
	db *sql.DB
}

func (BasicTradeStore) Create(trade *models.BasicTrade) error {
	panic("implement me")
}

func (BasicTradeStore) Get(ids ...int) ([]*models.BasicTrade, error) {
	panic("implement me")
}

func (BasicTradeStore) Update(trade *models.BasicTrade) error {
	panic("implement me")
}

func (BasicTradeStore) Delete(int) error {
	panic("implement me")
}

func NewBasicTradeStore(db *sql.DB) *BasicTradeStore {
	return &BasicTradeStore{
		db: db,
	}
}
