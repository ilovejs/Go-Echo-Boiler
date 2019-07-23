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

func testClaimHistories(t *testing.T) {
	t.Parallel()

	query := ClaimHistories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testClaimHistoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
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

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClaimHistoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ClaimHistories().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClaimHistoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClaimHistorySlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClaimHistoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ClaimHistoryExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ClaimHistory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ClaimHistoryExists to return true, but got false.")
	}
}

func testClaimHistoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	claimHistoryFound, err := FindClaimHistory(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if claimHistoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testClaimHistoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ClaimHistories().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testClaimHistoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ClaimHistories().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testClaimHistoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimHistoryOne := &ClaimHistory{}
	claimHistoryTwo := &ClaimHistory{}
	if err = randomize.Struct(seed, claimHistoryOne, claimHistoryDBTypes, false, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, claimHistoryTwo, claimHistoryDBTypes, false, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = claimHistoryOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = claimHistoryTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClaimHistories().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testClaimHistoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	claimHistoryOne := &ClaimHistory{}
	claimHistoryTwo := &ClaimHistory{}
	if err = randomize.Struct(seed, claimHistoryOne, claimHistoryDBTypes, false, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, claimHistoryTwo, claimHistoryDBTypes, false, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = claimHistoryOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = claimHistoryTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testClaimHistoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClaimHistoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(claimHistoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClaimHistoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
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

func testClaimHistoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClaimHistorySlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testClaimHistoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClaimHistories().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	claimHistoryDBTypes = map[string]string{`ID`: `int`, `TradeID`: `int`, `ClaimID`: `int`, `ProfileID`: `int`, `PreviousClaimed`: `float`, `IsActive`: `bit`, `IsDeleted`: `bit`, `Created`: `datetime`, `Updated`: `datetime`}
	_                   = bytes.MinRead
)

func testClaimHistoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(claimHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(claimHistoryAllColumns) == len(claimHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testClaimHistoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(claimHistoryAllColumns) == len(claimHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClaimHistory{}
	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, claimHistoryDBTypes, true, claimHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(claimHistoryAllColumns, claimHistoryPrimaryKeyColumns) {
		fields = claimHistoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			claimHistoryAllColumns,
			claimHistoryPrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(
			fields,
			claimHistoryColumnsWithAuto,
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

	slice := ClaimHistorySlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testClaimHistoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(claimHistoryAllColumns) == len(claimHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ClaimHistory{}
	if err = randomize.Struct(seed, &o, claimHistoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClaimHistory: %s", err)
	}

	count, err := ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, claimHistoryDBTypes, false, claimHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClaimHistory struct: %s", err)
	}

	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClaimHistory: %s", err)
	}

	count, err = ClaimHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
