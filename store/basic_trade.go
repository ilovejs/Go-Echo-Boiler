package store

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"onsite/models"
	"time"
)

type BasicTradeStore struct {
	db *sql.DB
}

func NewBasicTradeStore(db *sql.DB) *BasicTradeStore {
	return &BasicTradeStore{
		db: db,
	}
}

func (bts BasicTradeStore) Create(trade *models.BasicTrade) error {
	trade.Created.SetValid(time.Now())
	err := trade.Insert(bts.db, boil.Infer())
	if err != nil {
		return err
	}
	fmt.Println("Create basic trade")
	return nil
}

func (bts BasicTradeStore) Get(ids ...int) ([]*models.BasicTrade, error) {
	var bt []*models.BasicTrade
	var single *models.BasicTrade
	var err error

	if len(ids) == 1 {
		single, err = models.BasicTrades(Where("id = ? and is_deleted = ?", ids, false)).
			One(bts.db)
	} else {
		bt, err = models.BasicTrades(Where("id = ? and is_deleted = ?", ids, false)).
			All(bts.db)
	}

	if bt == nil {
		return nil, errors.New("basic trade not found")
	}
	if err != nil {
		return nil, err
	}
	//compile multiple
	var multiple = make([]*models.BasicTrade, 0)
	multiple = append(multiple, single)
	return multiple, nil
}

func (bts BasicTradeStore) Update(trade *models.BasicTrade) error {
	panic("implement me")
}

func (bts BasicTradeStore) Delete(int) error {
	panic("implement me")
}
