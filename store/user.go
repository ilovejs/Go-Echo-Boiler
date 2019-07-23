package store

import (
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	m "onsite/models"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) Create(u *m.User) (err error) {
	err = u.Insert(us.db, boil.Infer())
	if err != nil {
		return err
	}
	fmt.Println("Store: Create User", u.ID)
	return nil
}

func (us *UserStore) GetByID(id int) (*m.User, error) {
	u, err := m.Users(Where("id = ?", id)).One(us.db)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (us *UserStore) GetByEmail(e string) (*m.User, error) {
	u, err := m.Users(Where("email = ?", e)).One(us.db)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (us *UserStore) GetByUserName(u string) (*m.User, error) {
	panic("not done")
}

func (us *UserStore) Update(u *m.User) error {
	panic("not done")
}

func (us *UserStore) Delete(u *m.User) error {
	panic("not done")
}
