// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/types"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Project is an object representing the database table.
type Project struct {
	ID                   int64             `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID               int64             `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Name                 null.String       `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	TotalItemBreakdown   types.NullDecimal `boil:"total_item_breakdown" json:"total_item_breakdown,omitempty" toml:"total_item_breakdown" yaml:"total_item_breakdown,omitempty"`
	ContractorTotalClaim types.NullDecimal `boil:"contractor_total_claim" json:"contractor_total_claim,omitempty" toml:"contractor_total_claim" yaml:"contractor_total_claim,omitempty"`
	IsDeleted            null.Bool         `boil:"is_deleted" json:"is_deleted,omitempty" toml:"is_deleted" yaml:"is_deleted,omitempty"`
	Created              null.Time         `boil:"created" json:"created,omitempty" toml:"created" yaml:"created,omitempty"`
	Updated              null.Time         `boil:"updated" json:"updated,omitempty" toml:"updated" yaml:"updated,omitempty"`
	LoginID              int64             `boil:"login_id" json:"login_id" toml:"login_id" yaml:"login_id"`
	SerialNo             null.String       `boil:"serial_no" json:"serial_no,omitempty" toml:"serial_no" yaml:"serial_no,omitempty"`
	Details              null.String       `boil:"details" json:"details,omitempty" toml:"details" yaml:"details,omitempty"`
	TotalContractValue   types.NullDecimal `boil:"total_contract_value" json:"total_contract_value,omitempty" toml:"total_contract_value" yaml:"total_contract_value,omitempty"`
	QuantitySurveyor     null.String       `boil:"quantity_surveyor" json:"quantity_surveyor,omitempty" toml:"quantity_surveyor" yaml:"quantity_surveyor,omitempty"`
	Notes                null.String       `boil:"notes" json:"notes,omitempty" toml:"notes" yaml:"notes,omitempty"`

	R *projectR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L projectL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProjectColumns = struct {
	ID                   string
	UserID               string
	Name                 string
	TotalItemBreakdown   string
	ContractorTotalClaim string
	IsDeleted            string
	Created              string
	Updated              string
	LoginID              string
	SerialNo             string
	Details              string
	TotalContractValue   string
	QuantitySurveyor     string
	Notes                string
}{
	ID:                   "id",
	UserID:               "user_id",
	Name:                 "name",
	TotalItemBreakdown:   "total_item_breakdown",
	ContractorTotalClaim: "contractor_total_claim",
	IsDeleted:            "is_deleted",
	Created:              "created",
	Updated:              "updated",
	LoginID:              "login_id",
	SerialNo:             "serial_no",
	Details:              "details",
	TotalContractValue:   "total_contract_value",
	QuantitySurveyor:     "quantity_surveyor",
	Notes:                "notes",
}

// Generated where

var ProjectWhere = struct {
	ID                   whereHelperint64
	UserID               whereHelperint64
	Name                 whereHelpernull_String
	TotalItemBreakdown   whereHelpertypes_NullDecimal
	ContractorTotalClaim whereHelpertypes_NullDecimal
	IsDeleted            whereHelpernull_Bool
	Created              whereHelpernull_Time
	Updated              whereHelpernull_Time
	LoginID              whereHelperint64
	SerialNo             whereHelpernull_String
	Details              whereHelpernull_String
	TotalContractValue   whereHelpertypes_NullDecimal
	QuantitySurveyor     whereHelpernull_String
	Notes                whereHelpernull_String
}{
	ID:                   whereHelperint64{field: "[dbo].[projects].[id]"},
	UserID:               whereHelperint64{field: "[dbo].[projects].[user_id]"},
	Name:                 whereHelpernull_String{field: "[dbo].[projects].[name]"},
	TotalItemBreakdown:   whereHelpertypes_NullDecimal{field: "[dbo].[projects].[total_item_breakdown]"},
	ContractorTotalClaim: whereHelpertypes_NullDecimal{field: "[dbo].[projects].[contractor_total_claim]"},
	IsDeleted:            whereHelpernull_Bool{field: "[dbo].[projects].[is_deleted]"},
	Created:              whereHelpernull_Time{field: "[dbo].[projects].[created]"},
	Updated:              whereHelpernull_Time{field: "[dbo].[projects].[updated]"},
	LoginID:              whereHelperint64{field: "[dbo].[projects].[login_id]"},
	SerialNo:             whereHelpernull_String{field: "[dbo].[projects].[serial_no]"},
	Details:              whereHelpernull_String{field: "[dbo].[projects].[details]"},
	TotalContractValue:   whereHelpertypes_NullDecimal{field: "[dbo].[projects].[total_contract_value]"},
	QuantitySurveyor:     whereHelpernull_String{field: "[dbo].[projects].[quantity_surveyor]"},
	Notes:                whereHelpernull_String{field: "[dbo].[projects].[notes]"},
}

// ProjectRels is where relationship names are stored.
var ProjectRels = struct {
	Login              string
	User               string
	ContractorProjects string
	TradeItems         string
}{
	Login:              "Login",
	User:               "User",
	ContractorProjects: "ContractorProjects",
	TradeItems:         "TradeItems",
}

// projectR is where relationships are stored.
type projectR struct {
	Login              *Profile
	User               *Profile
	ContractorProjects ContractorProjectSlice
	TradeItems         TradeItemSlice
}

// NewStruct creates a new relationship struct
func (*projectR) NewStruct() *projectR {
	return &projectR{}
}

// projectL is where Load methods for each relationship are stored.
type projectL struct{}

var (
	projectAllColumns            = []string{"id", "user_id", "name", "total_item_breakdown", "contractor_total_claim", "is_deleted", "created", "updated", "login_id", "serial_no", "details", "total_contract_value", "quantity_surveyor", "notes"}
	projectColumnsWithAuto       = []string{}
	projectColumnsWithoutDefault = []string{"user_id", "name", "total_item_breakdown", "contractor_total_claim", "is_deleted", "created", "updated", "login_id", "serial_no", "details", "total_contract_value", "quantity_surveyor", "notes"}
	projectColumnsWithDefault    = []string{"id"}
	projectPrimaryKeyColumns     = []string{"id"}
)

type (
	// ProjectSlice is an alias for a slice of pointers to Project.
	// This should generally be used opposed to []Project.
	ProjectSlice []*Project

	projectQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	projectType                 = reflect.TypeOf(&Project{})
	projectMapping              = queries.MakeStructMapping(projectType)
	projectPrimaryKeyMapping, _ = queries.BindMapping(projectType, projectMapping, projectPrimaryKeyColumns)
	projectInsertCacheMut       sync.RWMutex
	projectInsertCache          = make(map[string]insertCache)
	projectUpdateCacheMut       sync.RWMutex
	projectUpdateCache          = make(map[string]updateCache)
	projectUpsertCacheMut       sync.RWMutex
	projectUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single project record from the query.
func (q projectQuery) One(exec boil.Executor) (*Project, error) {
	o := &Project{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for projects")
	}

	return o, nil
}

// All returns all Project records from the query.
func (q projectQuery) All(exec boil.Executor) (ProjectSlice, error) {
	var o []*Project

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Project slice")
	}

	return o, nil
}

// Count returns the count of all Project records in the query.
func (q projectQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count projects rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q projectQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if projects exists")
	}

	return count > 0, nil
}

// Login pointed to by the foreign key.
func (o *Project) Login(mods ...qm.QueryMod) profileQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.LoginID),
	}

	queryMods = append(queryMods, mods...)

	query := Profiles(queryMods...)
	queries.SetFrom(query.Query, "[dbo].[profiles]")

	return query
}

// User pointed to by the foreign key.
func (o *Project) User(mods ...qm.QueryMod) profileQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Profiles(queryMods...)
	queries.SetFrom(query.Query, "[dbo].[profiles]")

	return query
}

// ContractorProjects retrieves all the contractor_project's ContractorProjects with an executor.
func (o *Project) ContractorProjects(mods ...qm.QueryMod) contractorProjectQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("[dbo].[contractor_projects].[project_id]=?", o.ID),
	)

	query := ContractorProjects(queryMods...)
	queries.SetFrom(query.Query, "[dbo].[contractor_projects]")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"[dbo].[contractor_projects].*"})
	}

	return query
}

// TradeItems retrieves all the trade_item's TradeItems with an executor.
func (o *Project) TradeItems(mods ...qm.QueryMod) tradeItemQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("[dbo].[trade_items].[project_id]=?", o.ID),
	)

	query := TradeItems(queryMods...)
	queries.SetFrom(query.Query, "[dbo].[trade_items]")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"[dbo].[trade_items].*"})
	}

	return query
}

// LoadLogin allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (projectL) LoadLogin(e boil.Executor, singular bool, maybeProject interface{}, mods queries.Applicator) error {
	var slice []*Project
	var object *Project

	if singular {
		object = maybeProject.(*Project)
	} else {
		slice = *maybeProject.(*[]*Project)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectR{}
		}
		args = append(args, object.LoginID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectR{}
			}

			for _, a := range args {
				if a == obj.LoginID {
					continue Outer
				}
			}

			args = append(args, obj.LoginID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dbo.profiles`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Profile")
	}

	var resultSlice []*Profile
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Profile")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for profiles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for profiles")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Login = foreign
		if foreign.R == nil {
			foreign.R = &profileR{}
		}
		foreign.R.LoginProjects = append(foreign.R.LoginProjects, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LoginID == foreign.ID {
				local.R.Login = foreign
				if foreign.R == nil {
					foreign.R = &profileR{}
				}
				foreign.R.LoginProjects = append(foreign.R.LoginProjects, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (projectL) LoadUser(e boil.Executor, singular bool, maybeProject interface{}, mods queries.Applicator) error {
	var slice []*Project
	var object *Project

	if singular {
		object = maybeProject.(*Project)
	} else {
		slice = *maybeProject.(*[]*Project)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dbo.profiles`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Profile")
	}

	var resultSlice []*Profile
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Profile")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for profiles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for profiles")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &profileR{}
		}
		foreign.R.UserProjects = append(foreign.R.UserProjects, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &profileR{}
				}
				foreign.R.UserProjects = append(foreign.R.UserProjects, local)
				break
			}
		}
	}

	return nil
}

// LoadContractorProjects allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (projectL) LoadContractorProjects(e boil.Executor, singular bool, maybeProject interface{}, mods queries.Applicator) error {
	var slice []*Project
	var object *Project

	if singular {
		object = maybeProject.(*Project)
	} else {
		slice = *maybeProject.(*[]*Project)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dbo.contractor_projects`), qm.WhereIn(`project_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load contractor_projects")
	}

	var resultSlice []*ContractorProject
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice contractor_projects")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on contractor_projects")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for contractor_projects")
	}

	if singular {
		object.R.ContractorProjects = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &contractorProjectR{}
			}
			foreign.R.Project = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ProjectID {
				local.R.ContractorProjects = append(local.R.ContractorProjects, foreign)
				if foreign.R == nil {
					foreign.R = &contractorProjectR{}
				}
				foreign.R.Project = local
				break
			}
		}
	}

	return nil
}

// LoadTradeItems allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (projectL) LoadTradeItems(e boil.Executor, singular bool, maybeProject interface{}, mods queries.Applicator) error {
	var slice []*Project
	var object *Project

	if singular {
		object = maybeProject.(*Project)
	} else {
		slice = *maybeProject.(*[]*Project)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`dbo.trade_items`), qm.WhereIn(`project_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load trade_items")
	}

	var resultSlice []*TradeItem
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice trade_items")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on trade_items")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for trade_items")
	}

	if singular {
		object.R.TradeItems = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &tradeItemR{}
			}
			foreign.R.Project = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ProjectID {
				local.R.TradeItems = append(local.R.TradeItems, foreign)
				if foreign.R == nil {
					foreign.R = &tradeItemR{}
				}
				foreign.R.Project = local
				break
			}
		}
	}

	return nil
}

// SetLogin of the project to the related item.
// Sets o.R.Login to related.
// Adds o to related.R.LoginProjects.
func (o *Project) SetLogin(exec boil.Executor, insert bool, related *Profile) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE [dbo].[projects] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, []string{"login_id"}),
		strmangle.WhereClause("[", "]", 2, projectPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.LoginID = related.ID
	if o.R == nil {
		o.R = &projectR{
			Login: related,
		}
	} else {
		o.R.Login = related
	}

	if related.R == nil {
		related.R = &profileR{
			LoginProjects: ProjectSlice{o},
		}
	} else {
		related.R.LoginProjects = append(related.R.LoginProjects, o)
	}

	return nil
}

// SetUser of the project to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserProjects.
func (o *Project) SetUser(exec boil.Executor, insert bool, related *Profile) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE [dbo].[projects] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, []string{"user_id"}),
		strmangle.WhereClause("[", "]", 2, projectPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &projectR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &profileR{
			UserProjects: ProjectSlice{o},
		}
	} else {
		related.R.UserProjects = append(related.R.UserProjects, o)
	}

	return nil
}

// AddContractorProjects adds the given related objects to the existing relationships
// of the project, optionally inserting them as new records.
// Appends related to o.R.ContractorProjects.
// Sets related.R.Project appropriately.
func (o *Project) AddContractorProjects(exec boil.Executor, insert bool, related ...*ContractorProject) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ProjectID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE [dbo].[contractor_projects] SET %s WHERE %s",
				strmangle.SetParamNames("[", "]", 1, []string{"project_id"}),
				strmangle.WhereClause("[", "]", 2, contractorProjectPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ProjectID = o.ID
		}
	}

	if o.R == nil {
		o.R = &projectR{
			ContractorProjects: related,
		}
	} else {
		o.R.ContractorProjects = append(o.R.ContractorProjects, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &contractorProjectR{
				Project: o,
			}
		} else {
			rel.R.Project = o
		}
	}
	return nil
}

// AddTradeItems adds the given related objects to the existing relationships
// of the project, optionally inserting them as new records.
// Appends related to o.R.TradeItems.
// Sets related.R.Project appropriately.
func (o *Project) AddTradeItems(exec boil.Executor, insert bool, related ...*TradeItem) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ProjectID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE [dbo].[trade_items] SET %s WHERE %s",
				strmangle.SetParamNames("[", "]", 1, []string{"project_id"}),
				strmangle.WhereClause("[", "]", 2, tradeItemPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ProjectID = o.ID
		}
	}

	if o.R == nil {
		o.R = &projectR{
			TradeItems: related,
		}
	} else {
		o.R.TradeItems = append(o.R.TradeItems, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &tradeItemR{
				Project: o,
			}
		} else {
			rel.R.Project = o
		}
	}
	return nil
}

// Projects retrieves all the records using an executor.
func Projects(mods ...qm.QueryMod) projectQuery {
	mods = append(mods, qm.From("[dbo].[projects]"))
	return projectQuery{NewQuery(mods...)}
}

// FindProject retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProject(exec boil.Executor, iD int64, selectCols ...string) (*Project, error) {
	projectObj := &Project{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[projects] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, projectObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from projects")
	}

	return projectObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Project) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no projects provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(projectColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	projectInsertCacheMut.RLock()
	cache, cached := projectInsertCache[key]
	projectInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			projectAllColumns,
			projectColumnsWithDefault,
			projectColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(projectType, projectMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(projectType, projectMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[projects] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[projects] %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryOutput = fmt.Sprintf("OUTPUT INSERTED.[%s] ", strings.Join(returnColumns, "],INSERTED.["))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into projects")
	}

	if !cached {
		projectInsertCacheMut.Lock()
		projectInsertCache[key] = cache
		projectInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Project.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Project) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	projectUpdateCacheMut.RLock()
	cache, cached := projectUpdateCache[key]
	projectUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			projectAllColumns,
			projectPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, projectColumnsWithAuto)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update projects, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[projects] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, projectPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(projectType, projectMapping, append(wl, projectPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update projects row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for projects")
	}

	if !cached {
		projectUpdateCacheMut.Lock()
		projectUpdateCache[key] = cache
		projectUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q projectQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for projects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for projects")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProjectSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[projects] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, projectPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in project slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all project")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Project) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no projects provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(projectColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	projectUpsertCacheMut.RLock()
	cache, cached := projectUpsertCache[key]
	projectUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			projectAllColumns,
			projectColumnsWithDefault,
			projectColumnsWithoutDefault,
			nzDefaults,
		)
		insert = strmangle.SetComplement(insert, projectColumnsWithAuto)
		for i, v := range insert {
			if strmangle.ContainsAny(projectPrimaryKeyColumns, v) && strmangle.ContainsAny(projectColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("models: unable to upsert projects, could not build insert column list")
		}

		ret = strmangle.SetMerge(ret, projectColumnsWithAuto)
		ret = strmangle.SetMerge(ret, projectColumnsWithDefault)

		update := updateColumns.UpdateColumnSet(
			projectAllColumns,
			projectPrimaryKeyColumns,
		)
		update = strmangle.SetComplement(update, projectColumnsWithAuto)

		if len(update) == 0 {
			return errors.New("models: unable to upsert projects, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "projects", projectPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(projectPrimaryKeyColumns))
		copy(whitelist, projectPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(projectType, projectMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(projectType, projectMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // MSSQL doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert projects")
	}

	if !cached {
		projectUpsertCacheMut.Lock()
		projectUpsertCache[key] = cache
		projectUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Project record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Project) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Project provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), projectPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[projects] WHERE [id]=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from projects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for projects")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q projectQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no projectQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from projects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for projects")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProjectSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[projects] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, projectPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from project slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for projects")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Project) Reload(exec boil.Executor) error {
	ret, err := FindProject(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProjectSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProjectSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[projects].* FROM [dbo].[projects] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, projectPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProjectSlice")
	}

	*o = slice

	return nil
}

// ProjectExists checks if the Project row exists.
func ProjectExists(exec boil.Executor, iD int64) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[projects] where [id]=$1) then 1 else 0 end"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if projects exists")
	}

	return exists, nil
}
