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

func (ps ProjectStore) CreateProject(p *model.Project) (*model.Project, error) {
	err := ps.db.Debug().FirstOrCreate(
		&p,
		&model.Project{Name: p.Name}).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (ps ProjectStore) ReadProject(id int) (*model.Project, error) {
	var m model.Project

	if err := ps.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

//GetAll: Active Contractors
func (ps ProjectStore) GetContractors() ([]model.User, error) {
	var users []model.User

	err := ps.db.Debug().Model(&model.User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, err
}
