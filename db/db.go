package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"onsite/model"
	"os"
)

func New() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./realworld.db")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func TestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../realworld_test.db")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

func DropTestDB() error {
	if err := os.Remove("./../realworld_test.db"); err != nil {
		return err
	}
	return nil
}

//register table to migrate
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Follow{},
		&model.Article{},
		&model.Comment{},
		&model.Tag{},
		//
		&model.Role{},
		&model.Login{},
		//&model.UserProfile{},
		&model.Project{},
		&model.ContractorProject{},
		//&model.Item{},
		//&model.ItemType{},
		//&model.Claim{},
		//&model.ClaimHistory{},
	)
}
