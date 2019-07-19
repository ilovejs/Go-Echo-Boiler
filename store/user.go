package store

import (
	"database/sql"
	"onsite/models"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) Create(u *models.User) (err error) {
	panic("x")
}

func (us *UserStore) GetByID(id uint) (*models.User, error) {

	panic("not done")
}

func (us *UserStore) GetByEmail(e string) (*models.User, error) {

	panic("not done")
}

func (us *UserStore) GetByUserUserName(u string) (*models.User, error) {
	panic("not done")
}

func (us *UserStore) Update(u *models.User) error {
	panic("not done")
}

func (us *UserStore) Delete(u *models.User) error {
	panic("not done")
}
