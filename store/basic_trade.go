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
		single, err = models.BasicTrades(Where("id = ? and is_deleted = ?", ids[0], false)).One(bts.db)
		if single == nil {
			return nil, errors.New("basic trade not found")
		}
	} else {
		//todo: test not covered
		bt, err = models.BasicTrades(WhereIn("id in ?", ids)).All(bts.db)
		if bt == nil {
			return nil, errors.New("basic trade not found")
		}
	}

	if err != nil {
		return nil, err
	}
	//compile multiple
	var multiple = make([]*models.BasicTrade, 0)
	multiple = append(multiple, single)
	return multiple, nil
}

func (bts BasicTradeStore) List(offset int, limit int) ([]*models.BasicTrade, int, error) {
	AllBts, err := models.BasicTrades(Where("is_deleted = ?", false),
		Limit(limit),
		Offset(offset),
	).All(bts.db)

	if err != nil {
		return nil, 0, err
	}
	return AllBts, len(AllBts), nil
}

func (bts BasicTradeStore) Update(trade *models.BasicTrade) error {
	panic("implement me")
}

func (bts BasicTradeStore) Delete(id int) error {
	arr, err := bts.Get(id)
	if len(arr) != 1 {
		return errors.New("item not found, wrong id for basic trade")
	}
	if err != nil {
		//no record
		return err
	}
	basicTradeToDelete := arr[0]
	basicTradeToDelete.IsDeleted = true
	rowAffected, err := basicTradeToDelete.Update(bts.db, boil.Whitelist("is_deleted"))
	if err != nil {
		return err
	}
	fmt.Println("Delete BasicTrade row affect: ", rowAffected)
	return nil
}
