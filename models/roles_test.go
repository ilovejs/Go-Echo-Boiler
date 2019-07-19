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

func testRoles(t *testing.T) {
	t.Parallel()

	query := Roles()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRolesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
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

	count, err := Roles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRolesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Roles().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Roles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRolesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RoleSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Roles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRolesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RoleExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Role exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RoleExists to return true, but got false.")
	}
}

func testRolesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	roleFound, err := FindRole(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if roleFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRolesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Roles().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testRolesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Roles().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRolesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	roleOne := &Role{}
	roleTwo := &Role{}
	if err = randomize.Struct(seed, roleOne, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}
	if err = randomize.Struct(seed, roleTwo, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = roleOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = roleTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Roles().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRolesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	roleOne := &Role{}
	roleTwo := &Role{}
	if err = randomize.Struct(seed, roleOne, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}
	if err = randomize.Struct(seed, roleTwo, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = roleOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = roleTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Roles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testRolesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Roles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRolesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(roleColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Roles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRoleToManyUserRoleLogins(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Role
	var b, c Login

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, loginDBTypes, false, loginColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, loginDBTypes, false, loginColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.UserRoleID, a.ID)
	queries.Assign(&c.UserRoleID, a.ID)
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserRoleLogins().All(tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.UserRoleID, b.UserRoleID) {
			bFound = true
		}
		if queries.Equal(v.UserRoleID, c.UserRoleID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := RoleSlice{&a}
	if err = a.L.LoadUserRoleLogins(tx, false, (*[]*Role)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserRoleLogins); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserRoleLogins = nil
	if err = a.L.LoadUserRoleLogins(tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserRoleLogins); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testRoleToManyAddOpUserRoleLogins(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Role
	var b, c, d, e Login

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Login{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, loginDBTypes, false, strmangle.SetComplement(loginPrimaryKeyColumns, loginColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Login{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserRoleLogins(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.UserRoleID) {
			t.Error("foreign key was wrong value", a.ID, first.UserRoleID)
		}
		if !queries.Equal(a.ID, second.UserRoleID) {
			t.Error("foreign key was wrong value", a.ID, second.UserRoleID)
		}

		if first.R.UserRole != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.UserRole != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserRoleLogins[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserRoleLogins[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserRoleLogins().Count(tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testRoleToManySetOpUserRoleLogins(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Role
	var b, c, d, e Login

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Login{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, loginDBTypes, false, strmangle.SetComplement(loginPrimaryKeyColumns, loginColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetUserRoleLogins(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.UserRoleLogins().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetUserRoleLogins(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.UserRoleLogins().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.UserRoleID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.UserRoleID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.UserRoleID) {
		t.Error("foreign key was wrong value", a.ID, d.UserRoleID)
	}
	if !queries.Equal(a.ID, e.UserRoleID) {
		t.Error("foreign key was wrong value", a.ID, e.UserRoleID)
	}

	if b.R.UserRole != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.UserRole != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.UserRole != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.UserRole != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.UserRoleLogins[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.UserRoleLogins[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testRoleToManyRemoveOpUserRoleLogins(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()

	var a Role
	var b, c, d, e Login

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Login{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, loginDBTypes, false, strmangle.SetComplement(loginPrimaryKeyColumns, loginColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddUserRoleLogins(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.UserRoleLogins().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveUserRoleLogins(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.UserRoleLogins().Count(tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.UserRoleID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.UserRoleID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.UserRole != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.UserRole != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.UserRole != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.UserRole != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.UserRoleLogins) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.UserRoleLogins[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.UserRoleLogins[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testRolesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
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

func testRolesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RoleSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testRolesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Roles().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	roleDBTypes = map[string]string{`ID`: `int`, `UserRoleType`: `varchar`, `Created`: `datetime`}
	_           = bytes.MinRead
)

func testRolesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(rolePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(roleAllColumns) == len(rolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Roles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, roleDBTypes, true, rolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRolesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(roleAllColumns) == len(rolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Roles().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, roleDBTypes, true, rolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(roleAllColumns, rolePrimaryKeyColumns) {
		fields = roleAllColumns
	} else {
		fields = strmangle.SetComplement(
			roleAllColumns,
			rolePrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(
			fields,
			roleColumnsWithAuto,
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

	slice := RoleSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRolesUpsert(t *testing.T) {
	t.Parallel()

	if len(roleAllColumns) == len(rolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Role{}
	if err = randomize.Struct(seed, &o, roleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Role: %s", err)
	}

	count, err := Roles().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, roleDBTypes, false, rolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Role: %s", err)
	}

	count, err = Roles().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}