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

func testProfiles(t *testing.T) {
	t.Parallel()

	query := Profiles()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProfilesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
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

	count, err := Profiles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProfilesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Profiles().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Profiles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProfilesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProfileSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Profiles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProfilesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProfileExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Profile exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProfileExists to return true, but got false.")
	}
}

func testProfilesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	profileFound, err := FindProfile(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if profileFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProfilesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Profiles().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testProfilesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Profiles().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProfilesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	profileOne := &Profile{}
	profileTwo := &Profile{}
	if err = randomize.Struct(seed, profileOne, profileDBTypes, false, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}
	if err = randomize.Struct(seed, profileTwo, profileDBTypes, false, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = profileOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = profileTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Profiles().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProfilesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	profileOne := &Profile{}
	profileTwo := &Profile{}
	if err = randomize.Struct(seed, profileOne, profileDBTypes, false, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}
	if err = randomize.Struct(seed, profileTwo, profileDBTypes, false, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = profileOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = profileTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Profiles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testProfilesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Profiles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProfilesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(profileColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Profiles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProfileToManyLoginClaimHistories(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c ClaimHistory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, claimHistoryDBTypes, false, claimHistoryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, claimHistoryDBTypes, false, claimHistoryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.LoginID = a.ID
	c.LoginID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.LoginClaimHistories().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.LoginID == b.LoginID {
			bFound = true
		}
		if v.LoginID == c.LoginID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ProfileSlice{&a}
	if err = a.L.LoadLoginClaimHistories(tx, false, (*[]*Profile)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LoginClaimHistories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.LoginClaimHistories = nil
	if err = a.L.LoadLoginClaimHistories(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LoginClaimHistories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testProfileToManyLoginClaims(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
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

	b.LoginID = a.ID
	c.LoginID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.LoginClaims().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.LoginID == b.LoginID {
			bFound = true
		}
		if v.LoginID == c.LoginID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ProfileSlice{&a}
	if err = a.L.LoadLoginClaims(tx, false, (*[]*Profile)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LoginClaims); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.LoginClaims = nil
	if err = a.L.LoadLoginClaims(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LoginClaims); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testProfileToManyUserContractorProjects(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c ContractorProject

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
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

	b.UserID = a.ID
	c.UserID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserContractorProjects().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.UserID == b.UserID {
			bFound = true
		}
		if v.UserID == c.UserID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ProfileSlice{&a}
	if err = a.L.LoadUserContractorProjects(tx, false, (*[]*Profile)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserContractorProjects); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserContractorProjects = nil
	if err = a.L.LoadUserContractorProjects(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserContractorProjects); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testProfileToManyLoginProjects(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.LoginID = a.ID
	c.LoginID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.LoginProjects().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.LoginID == b.LoginID {
			bFound = true
		}
		if v.LoginID == c.LoginID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ProfileSlice{&a}
	if err = a.L.LoadLoginProjects(tx, false, (*[]*Profile)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LoginProjects); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.LoginProjects = nil
	if err = a.L.LoadLoginProjects(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LoginProjects); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testProfileToManyUserProjects(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.UserID = a.ID
	c.UserID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserProjects().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.UserID == b.UserID {
			bFound = true
		}
		if v.UserID == c.UserID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ProfileSlice{&a}
	if err = a.L.LoadUserProjects(tx, false, (*[]*Profile)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserProjects); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserProjects = nil
	if err = a.L.LoadUserProjects(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserProjects); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testProfileToManyLoginTradeItems(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c TradeItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
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

	b.LoginID = a.ID
	c.LoginID = a.ID

	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.LoginTradeItems().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.LoginID == b.LoginID {
			bFound = true
		}
		if v.LoginID == c.LoginID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ProfileSlice{&a}
	if err = a.L.LoadLoginTradeItems(tx, false, (*[]*Profile)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LoginTradeItems); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.LoginTradeItems = nil
	if err = a.L.LoadLoginTradeItems(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LoginTradeItems); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testProfileToManyAddOpLoginClaimHistories(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c, d, e ClaimHistory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, false, strmangle.SetComplement(profilePrimaryKeyColumns, profileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ClaimHistory{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, claimHistoryDBTypes, false, strmangle.SetComplement(claimHistoryPrimaryKeyColumns, claimHistoryColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*ClaimHistory{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddLoginClaimHistories(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.LoginID {
			t.Error("foreign key was wrong value", a.ID, first.LoginID)
		}
		if a.ID != second.LoginID {
			t.Error("foreign key was wrong value", a.ID, second.LoginID)
		}

		if first.R.Login != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Login != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.LoginClaimHistories[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.LoginClaimHistories[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.LoginClaimHistories().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testProfileToManyAddOpLoginClaims(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c, d, e Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, false, strmangle.SetComplement(profilePrimaryKeyColumns, profileColumnsWithoutDefault)...); err != nil {
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
		err = a.AddLoginClaims(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.LoginID {
			t.Error("foreign key was wrong value", a.ID, first.LoginID)
		}
		if a.ID != second.LoginID {
			t.Error("foreign key was wrong value", a.ID, second.LoginID)
		}

		if first.R.Login != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Login != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.LoginClaims[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.LoginClaims[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.LoginClaims().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testProfileToManyAddOpUserContractorProjects(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c, d, e ContractorProject

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, false, strmangle.SetComplement(profilePrimaryKeyColumns, profileColumnsWithoutDefault)...); err != nil {
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
		err = a.AddUserContractorProjects(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.UserID {
			t.Error("foreign key was wrong value", a.ID, first.UserID)
		}
		if a.ID != second.UserID {
			t.Error("foreign key was wrong value", a.ID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserContractorProjects[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserContractorProjects[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserContractorProjects().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testProfileToManyAddOpLoginProjects(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c, d, e Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, false, strmangle.SetComplement(profilePrimaryKeyColumns, profileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Project{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Project{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddLoginProjects(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.LoginID {
			t.Error("foreign key was wrong value", a.ID, first.LoginID)
		}
		if a.ID != second.LoginID {
			t.Error("foreign key was wrong value", a.ID, second.LoginID)
		}

		if first.R.Login != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Login != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.LoginProjects[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.LoginProjects[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.LoginProjects().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testProfileToManyAddOpUserProjects(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c, d, e Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, false, strmangle.SetComplement(profilePrimaryKeyColumns, profileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Project{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Project{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserProjects(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.UserID {
			t.Error("foreign key was wrong value", a.ID, first.UserID)
		}
		if a.ID != second.UserID {
			t.Error("foreign key was wrong value", a.ID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserProjects[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserProjects[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserProjects().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testProfileToManyAddOpLoginTradeItems(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c, d, e TradeItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, false, strmangle.SetComplement(profilePrimaryKeyColumns, profileColumnsWithoutDefault)...); err != nil {
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
		err = a.AddLoginTradeItems(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.LoginID {
			t.Error("foreign key was wrong value", a.ID, first.LoginID)
		}
		if a.ID != second.LoginID {
			t.Error("foreign key was wrong value", a.ID, second.LoginID)
		}

		if first.R.Login != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Login != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.LoginTradeItems[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.LoginTradeItems[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.LoginTradeItems().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testProfileToOneLoginUsingLogin(t *testing.T) {

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var local Profile
	var foreign Login

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, loginDBTypes, false, loginColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Login struct: %s", err)
	}

	if err := foreign.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.LoginID, foreign.ID)
	if err := local.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Login().One(tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ProfileSlice{&local}
	if err = local.L.LoadLogin(tx, false, (*[]*Profile)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Login == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Login = nil
	if err = local.L.LoadLogin(tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Login == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testProfileToOneSetOpLoginUsingLogin(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b, c Login

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, false, strmangle.SetComplement(profilePrimaryKeyColumns, profileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, loginDBTypes, false, strmangle.SetComplement(loginPrimaryKeyColumns, loginColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, loginDBTypes, false, strmangle.SetComplement(loginPrimaryKeyColumns, loginColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Login{&b, &c} {
		err = a.SetLogin(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Login != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Profiles[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.LoginID, x.ID) {
			t.Error("foreign key was wrong value", a.LoginID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.LoginID))
		reflect.Indirect(reflect.ValueOf(&a.LoginID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.LoginID, x.ID) {
			t.Error("foreign key was wrong value", a.LoginID, x.ID)
		}
	}
}

func testProfileToOneRemoveOpLoginUsingLogin(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Profile
	var b Login

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, profileDBTypes, false, strmangle.SetComplement(profilePrimaryKeyColumns, profileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, loginDBTypes, false, strmangle.SetComplement(loginPrimaryKeyColumns, loginColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetLogin(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveLogin(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Login().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Login != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.LoginID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Profiles) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testProfilesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
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

func testProfilesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProfileSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testProfilesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Profiles().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	profileDBTypes = map[string]string{`ID`: `bigint`, `LoginID`: `bigint`, `Mobile`: `varchar`, `FirstName`: `nvarchar`, `LastName`: `nvarchar`, `Email`: `nvarchar`, `IsActive`: `bit`, `IsDeleted`: `bit`, `Created`: `datetime`, `Updated`: `datetime`, `ForgetPasswordToken`: `varchar`, `TokenExpires`: `datetime`, `Company`: `nvarchar`}
	_              = bytes.MinRead
)

func testProfilesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(profilePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(profileAllColumns) == len(profilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Profiles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, profileDBTypes, true, profilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProfilesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(profileAllColumns) == len(profilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Profile{}
	if err = randomize.Struct(seed, o, profileDBTypes, true, profileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Profiles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, profileDBTypes, true, profilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(profileAllColumns, profilePrimaryKeyColumns) {
		fields = profileAllColumns
	} else {
		fields = strmangle.SetComplement(
			profileAllColumns,
			profilePrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(
			fields,
			profileColumnsWithAuto,
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

	slice := ProfileSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testProfilesUpsert(t *testing.T) {
	t.Parallel()

	if len(profileAllColumns) == len(profilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Profile{}
	if err = randomize.Struct(seed, &o, profileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Profile: %s", err)
	}

	count, err := Profiles().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, profileDBTypes, false, profilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Profile struct: %s", err)
	}

	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Profile: %s", err)
	}

	count, err = Profiles().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}