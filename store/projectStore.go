package store

import (
	"github.com/jinzhu/gorm"
	"onsite/model"
	_ "onsite/model"
)

type ProjectStore struct {
	db *gorm.DB
}

func NewProjectStore(db *gorm.DB) *ProjectStore {
	return &ProjectStore{
		db: db,
	}
}

func (ps ProjectStore) GetById(id uint) (*model.Project, error) {
	var m model.Project

	if err := ps.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (ps ProjectStore) CreateProject(p *model.Project) error {
	panic("not done")
}

//Active Contractors
func (ps ProjectStore) GetContractors() ([]model.User, error) {
	var users []model.User

	err := ps.db.Debug().Model(&model.User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, err
}
