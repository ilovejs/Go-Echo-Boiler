package store

import (
	"database/sql"
)

type ProjectStore struct {
	db *sql.DB
}

func NewProjectStore(db *sql.DB) *ProjectStore {
	return &ProjectStore{
		db: db,
	}
}

//func (ps ProjectStore) CreateProject(p *models.Project) (*models.Project, error) {
//}
//
//func (ps ProjectStore) ReadProject(id int) (*models.Project, error) {
//	var m models.Project
//
//}

//GetAll: Active Contractors
//func (ps ProjectStore) GetContractors() ([]models.User, error) {
//	var users []models.User
//
//	err := ps.db.Debug().models(&models.User{}).Find(&users).Error
//	if err != nil {
//		return nil, err
//	}
//	return users, err
//}
