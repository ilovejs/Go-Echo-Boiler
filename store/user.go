package store

import (
	"database/sql"
	"onsite/models"
)

type UserLoginStore struct {
	db *sql.DB
}

func NewUserLoginStore(db *sql.DB) *UserLoginStore {
	return &UserLoginStore{
		db: db,
	}
}

func (us *UserLoginStore) Create(u *models.Login) (err error) {
	panic("x")
}

func (us *UserLoginStore) GetByID(id uint) (*models.Login, error) {

	panic("not done")
}

func (us *UserLoginStore) GetByEmail(e string) (*models.Login, error) {

	panic("not done")
}

func (us *UserLoginStore) GetByUserLoginName(u string) (*models.Login, error) {
	panic("not done")
}

func (us *UserLoginStore) Update(u *models.Login) error {
	panic("not done")
}

func (us *UserLoginStore) Delete(u *models.Login) error {
	panic("not done")
}
