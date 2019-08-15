package store

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"onsite/models"
	"time"
)

type TradeStore struct {
	db *sql.DB
}

func NewTradeStore(db *sql.DB) *TradeStore {
	return &TradeStore{
		db: db,
	}
}

func (ts TradeStore) Create(tradeC *models.Trade) error {
	tradeC.Created.SetValid(time.Now())
	err := tradeC.Insert(ts.db, boil.Infer())
	if err != nil {
		return err
	}
	fmt.Println("Create trade")
	return nil
}

func (ts TradeStore) Get(ids ...int) ([]*models.Trade, error) {
	var bt []*models.Trade
	var single *models.Trade
	var err error

	if len(ids) == 1 {
		single, err = models.Trades(
			Load(models.TradeRels.TradeCategory),
			Where("id = ? and is_deleted = ?", ids[0], 0),
		).One(ts.db)

		if single == nil {
			return nil, errors.New("trade not found")
		}
	} else {
		// todo: test not covered
		bt, err = models.Trades(
			Load(models.TradeRels.TradeCategory),
			WhereIn("id in ?", ids),
		).All(ts.db)
		if bt == nil {
			return nil, errors.New("trade not found")
		}
	}

	if err != nil {
		return nil, err
	}

	var trades = make([]*models.Trade, 0)
	trades = append(trades, single)
	return trades, nil
}

func (ts TradeStore) List(offset int, limit int) ([]*models.Trade, int, error) {
	trades, err := models.Trades(
		Load(models.TradeRels.TradeCategory),
		Where("is_deleted = ?", 0),
		Offset(offset),
		Limit(limit),
	).All(ts.db)

	if err != nil {
		return nil, 0, err
	}
	return trades, len(trades), nil
}

func (ts TradeStore) Update(t *models.Trade) error {
	t.Updated.SetValid(time.Now())
	updated, err := t.Update(ts.db, boil.Infer())
	if err != nil {
		return err
	}
	fmt.Println("Trade Updated")
	spew.Dump(updated)
	return nil
}

func (ts TradeStore) Delete(id int) error {
	arr, err := ts.Get(id)
	if len(arr) != 1 {
		return errors.New("item not found, wrong id for trade")
	}
	if err != nil {
		// no record
		return err
	}
	TradeToDelete := arr[0]
	TradeToDelete.IsDeleted = true

	rowAffected, err := TradeToDelete.Update(
		ts.db,
		boil.Whitelist("is_deleted"),
	)
	if err != nil {
		return err
	}
	fmt.Println("Delete Trade row affect: ", rowAffected)
	return nil
}

func (ts TradeStore) GetByProject(projectId int) ([]*models.Trade, error) {
	panic("no")
}

func (ts TradeStore) Count(tradeId int) (int, error) {
	panic("not yet")
}
