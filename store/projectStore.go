package store

import (
	"github.com/jinzhu/gorm"
	"onsite/model"
	_ "onsite/model"
)

type ProjectStore struct {
	db *gorm.DB
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

func NewProjectStore(db *gorm.DB) *ProjectStore {
	return &ProjectStore{
		db: db,
	}
}
