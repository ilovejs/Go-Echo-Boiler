package store

import (
	"github.com/jinzhu/gorm"
	"onsite/model"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) Get(id uint) (*model.User, error) {
	var m model.User
	if err := us.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) Create(u *model.User) (err error) {
	return us.db.Create(u).Error
}

func (us *UserStore) Update(u *model.User) error {
	return us.db.Model(u).Update(u).Error
}

func (us *UserStore) Delete(u *model.User) error {
	return us.db.Model(u).Where(&model.User{Email: u.Email}).
		Update(&model.User{IsActive: false, IsDeleted: true}).Error
}
