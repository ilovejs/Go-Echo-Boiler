// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"github.com/volatiletech/sqlboiler/drivers"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

var dialect = drivers.Dialect{
	LQ: 0x5b,
	RQ: 0x5d,

	UseIndexPlaceholders:    true,
	UseLastInsertID:         false,
	UseSchema:               true,
	UseDefaultKeyword:       true,
	UseAutoColumns:          true,
	UseTopClause:            true,
	UseOutputClause:         true,
	UseCaseWhenExistsClause: true,
}

// NewQuery initializes a new Query using the passed in QueryMods
func NewQuery(mods ...qm.QueryMod) *queries.Query {
	q := &queries.Query{}
	queries.SetDialect(q, &dialect)
	qm.Apply(q, mods...)

	return q
}
