// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testBasicTradeitems(t *testing.T) {
	t.Parallel()

	query := BasicTradeitems()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBasicTradeitemsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BasicTradeitem{}
	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BasicTradeitems().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBasicTradeitemsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BasicTradeitem{}
	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BasicTradeitems().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BasicTradeitems().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBasicTradeitemsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BasicTradeitem{}
	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BasicTradeitemSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BasicTradeitems().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBasicTradeitemsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BasicTradeitem{}
	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BasicTradeitemExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BasicTradeitem exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BasicTradeitemExists to return true, but got false.")
	}
}

func testBasicTradeitemsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BasicTradeitem{}
	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	basicTradeitemFound, err := FindBasicTradeitem(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if basicTradeitemFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBasicTradeitemsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BasicTradeitem{}
	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BasicTradeitems().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testBasicTradeitemsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BasicTradeitem{}
	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BasicTradeitems().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBasicTradeitemsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	basicTradeitemOne := &BasicTradeitem{}
	basicTradeitemTwo := &BasicTradeitem{}
	if err = randomize.Struct(seed, basicTradeitemOne, basicTradeitemDBTypes, false, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}
	if err = randomize.Struct(seed, basicTradeitemTwo, basicTradeitemDBTypes, false, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = basicTradeitemOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = basicTradeitemTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BasicTradeitems().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBasicTradeitemsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	basicTradeitemOne := &BasicTradeitem{}
	basicTradeitemTwo := &BasicTradeitem{}
	if err = randomize.Struct(seed, basicTradeitemOne, basicTradeitemDBTypes, false, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}
	if err = randomize.Struct(seed, basicTradeitemTwo, basicTradeitemDBTypes, false, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = basicTradeitemOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = basicTradeitemTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BasicTradeitems().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testBasicTradeitemsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BasicTradeitem{}
	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BasicTradeitems().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBasicTradeitemsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BasicTradeitem{}
	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(basicTradeitemColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BasicTradeitems().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBasicTradeitemToManyTradeClaims(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a BasicTradeitem
	var b, c Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, claimDBTypes, false, claimColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, claimDBTypes, false, claimColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.TradeID = a.ID
	c.TradeID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.TradeClaims().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.TradeID == b.TradeID {
			bFound = true
		}
		if v.TradeID == c.TradeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BasicTradeitemSlice{&a}
	if err = a.L.LoadTradeClaims(tx, false, (*[]*BasicTradeitem)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TradeClaims); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TradeClaims = nil
	if err = a.L.LoadTradeClaims(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TradeClaims); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBasicTradeitemToManyTradeContractorProjects(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a BasicTradeitem
	var b, c ContractorProject

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, contractorProjectDBTypes, false, contractorProjectColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, contractorProjectDBTypes, false, contractorProjectColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.TradeID = a.ID
	c.TradeID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.TradeContractorProjects().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.TradeID == b.TradeID {
			bFound = true
		}
		if v.TradeID == c.TradeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BasicTradeitemSlice{&a}
	if err = a.L.LoadTradeContractorProjects(tx, false, (*[]*BasicTradeitem)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TradeContractorProjects); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TradeContractorProjects = nil
	if err = a.L.LoadTradeContractorProjects(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TradeContractorProjects); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBasicTradeitemToManyTradeTradeItems(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a BasicTradeitem
	var b, c TradeItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, tradeItemDBTypes, false, tradeItemColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, tradeItemDBTypes, false, tradeItemColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.TradeID = a.ID
	c.TradeID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.TradeTradeItems().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.TradeID == b.TradeID {
			bFound = true
		}
		if v.TradeID == c.TradeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BasicTradeitemSlice{&a}
	if err = a.L.LoadTradeTradeItems(tx, false, (*[]*BasicTradeitem)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TradeTradeItems); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TradeTradeItems = nil
	if err = a.L.LoadTradeTradeItems(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TradeTradeItems); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBasicTradeitemToManyAddOpTradeClaims(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a BasicTradeitem
	var b, c, d, e Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, basicTradeitemDBTypes, false, strmangle.SetComplement(basicTradeitemPrimaryKeyColumns, basicTradeitemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Claim{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Claim{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTradeClaims(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.TradeID {
			t.Error("foreign key was wrong value", a.ID, first.TradeID)
		}
		if a.ID != second.TradeID {
			t.Error("foreign key was wrong value", a.ID, second.TradeID)
		}

		if first.R.Trade != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Trade != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TradeClaims[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TradeClaims[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TradeClaims().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBasicTradeitemToManyAddOpTradeContractorProjects(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a BasicTradeitem
	var b, c, d, e ContractorProject

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, basicTradeitemDBTypes, false, strmangle.SetComplement(basicTradeitemPrimaryKeyColumns, basicTradeitemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ContractorProject{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, contractorProjectDBTypes, false, strmangle.SetComplement(contractorProjectPrimaryKeyColumns, contractorProjectColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ContractorProject{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTradeContractorProjects(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.TradeID {
			t.Error("foreign key was wrong value", a.ID, first.TradeID)
		}
		if a.ID != second.TradeID {
			t.Error("foreign key was wrong value", a.ID, second.TradeID)
		}

		if first.R.Trade != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Trade != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TradeContractorProjects[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TradeContractorProjects[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TradeContractorProjects().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBasicTradeitemToManyAddOpTradeTradeItems(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a BasicTradeitem
	var b, c, d, e TradeItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, basicTradeitemDBTypes, false, strmangle.SetComplement(basicTradeitemPrimaryKeyColumns, basicTradeitemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*TradeItem{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, tradeItemDBTypes, false, strmangle.SetComplement(tradeItemPrimaryKeyColumns, tradeItemColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*TradeItem{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTradeTradeItems(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.TradeID {
			t.Error("foreign key was wrong value", a.ID, first.TradeID)
		}
		if a.ID != second.TradeID {
			t.Error("foreign key was wrong value", a.ID, second.TradeID)
		}

		if first.R.Trade != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Trade != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TradeTradeItems[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TradeTradeItems[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TradeTradeItems().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBasicTradeitemsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BasicTradeitem{}
	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testBasicTradeitemsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BasicTradeitem{}
	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BasicTradeitemSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testBasicTradeitemsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BasicTradeitem{}
	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BasicTradeitems().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	basicTradeitemDBTypes = map[string]string{`ID`: `int`, `TradeName`: `varchar`, `CreatedOn`: `datetime`, `IsDeleted`: `bit`}
	_                     = bytes.MinRead
)

func testBasicTradeitemsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(basicTradeitemPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(basicTradeitemAllColumns) == len(basicTradeitemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BasicTradeitem{}
	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BasicTradeitems().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBasicTradeitemsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(basicTradeitemAllColumns) == len(basicTradeitemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BasicTradeitem{}
	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BasicTradeitems().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, basicTradeitemDBTypes, true, basicTradeitemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(basicTradeitemAllColumns, basicTradeitemPrimaryKeyColumns) {
		fields = basicTradeitemAllColumns
	} else {
		fields = strmangle.SetComplement(
			basicTradeitemAllColumns,
			basicTradeitemPrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(
			fields,
			basicTradeitemColumnsWithAuto,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := BasicTradeitemSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBasicTradeitemsUpsert(t *testing.T) {
	t.Parallel()

	if len(basicTradeitemAllColumns) == len(basicTradeitemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BasicTradeitem{}
	if err = randomize.Struct(seed, &o, basicTradeitemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BasicTradeitem: %s", err)
	}

	count, err := BasicTradeitems().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, basicTradeitemDBTypes, false, basicTradeitemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BasicTradeitem struct: %s", err)
	}

	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BasicTradeitem: %s", err)
	}

	count, err = BasicTradeitems().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}