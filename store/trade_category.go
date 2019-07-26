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

type TradeCategoryStore struct {
	db *sql.DB
}

func NewTradeCategoryStore(db *sql.DB) *TradeCategoryStore {
	return &TradeCategoryStore{
		db: db,
	}
}

func (bts TradeCategoryStore) Create(trade *models.TradeCategory) error {
	trade.Created.SetValid(time.Now())
	err := trade.Insert(bts.db, boil.Infer())
	if err != nil {
		return err
	}
	fmt.Println("Create trade catogory")
	return nil
}

func (bts TradeCategoryStore) Get(ids ...int) ([]*models.TradeCategory, error) {
	var bt []*models.TradeCategory
	var single *models.TradeCategory
	var err error

	if len(ids) == 1 {
		single, err = models.TradeCategories(Where("id = ? and is_deleted = ?", ids[0], false)).One(bts.db)
		if single == nil {
			return nil, errors.New("trade catogory not found")
		}
	} else {
		//todo: test not covered
		bt, err = models.TradeCategories(WhereIn("id in ?", ids)).All(bts.db)
		if bt == nil {
			return nil, errors.New("trade catogory not found")
		}
	}

	if err != nil {
		return nil, err
	}
	//compile multiple
	var multiple = make([]*models.TradeCategory, 0)
	multiple = append(multiple, single)
	return multiple, nil
}

func (bts TradeCategoryStore) List(offset int, limit int) ([]*models.TradeCategory, int, error) {
	AllBts, err := models.TradeCategories(Where("is_deleted = ?", false),
		Limit(limit),
		Offset(offset),
	).All(bts.db)

	if err != nil {
		return nil, 0, err
	}
	return AllBts, len(AllBts), nil
}

func (bts TradeCategoryStore) Update(t *models.TradeCategory) error {
	t.Updated.SetValid(time.Now())
	updated, err := t.Update(bts.db, boil.Infer())
	if err != nil {
		return err
	}
	fmt.Println("TradeCategory Updated")
	spew.Dump(updated)
	return nil
}

func (bts TradeCategoryStore) Delete(id int) error {
	arr, err := bts.Get(id)
	if len(arr) != 1 {
		return errors.New("item not found, wrong id for trade catogory")
	}
	if err != nil {
		//no record
		return err
	}
	TradeCategoryToDelete := arr[0]
	TradeCategoryToDelete.IsDeleted = true
	rowAffected, err := TradeCategoryToDelete.Update(bts.db, boil.Whitelist("is_deleted"))
	if err != nil {
		return err
	}
	fmt.Println("Delete TradeCategory row affect: ", rowAffected)
	return nil
}
