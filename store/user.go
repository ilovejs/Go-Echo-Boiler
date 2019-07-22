package store

import (
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/boil"
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
	newUser := &m.User{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password}

	err = newUser.Insert(us.db, boil.Infer())
	if err != nil {
		return err
	}
	fmt.Println("New User", newUser.ID)
	return nil
}

func (us *UserStore) GetByID(id uint) (*m.User, error) {

	panic("not done")
}

func (us *UserStore) GetByEmail(e string) (*m.User, error) {

	panic("not done")
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
