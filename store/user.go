package store

import (
	"database/sql"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	m "onsite/models"
	"time"
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

func (us *UserStore) GetByUserName(un string) (*m.User, error) {
	u, err := m.Users(Where("username = ?", un)).One(us.db)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (us *UserStore) Update(u *m.User) error {
	u.Updated.SetValid(time.Now())
	updated, err := u.Update(us.db, boil.Infer())
	if err != nil {
		return err
	}
	spew.Dump(updated)
	return nil
}

func (us *UserStore) List(offset int, limit int) ([]*m.User, int, error) {
	users, err := m.Users(Where("is_deleted = ?", 0),
		Limit(limit),
		Offset(offset),
	).All(us.db)

	if err != nil {
		return nil, 0, err
	}
	return users, len(users), nil
}

func (us *UserStore) Delete(uid int) error {
	panic("not done")
}
