package store

import (
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"onsite/models"
)

type ProjectStore struct {
	db *sql.DB
}

func NewProjectStore(db *sql.DB) *ProjectStore {
	return &ProjectStore{
		db: db,
	}
}

func (ps *ProjectStore) Create(p *models.Project) error {
	err := p.Insert(ps.db, boil.Infer())
	if err != nil {
		return err
	}
	fmt.Println("Create Project")
	return nil
}

func (ps *ProjectStore) Read(id int) (*models.Project, error) {
	p, err := models.Projects(Where("id = ?", id)).One(ps.db)
	if err != nil {
		return nil, err
	}
	return p, nil
}

//GetAll: Active Contractors
func (ps *ProjectStore) Contractors(pid int) ([]models.User, error) {
	panic("no")
	//var users []models.User

	//err := ps.db.Debug().models(&models.User{}).Find(&users).Error
	//if err != nil {
	//	return nil, err
	//}
	//return users, err
}
