package store

import "database/sql"

type ClaimStore struct {
	db *sql.DB
}

func NewClaimStore(db *sql.DB) *ClaimStore {
	return &ClaimStore{
		db: db,
	}
}
