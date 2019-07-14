package model

//
//import (
//	"github.com/jinzhu/gorm"
//)
//
//type Article struct {
//	gorm.Model
//	Slug        string `gorm:"unique_index;not null"`
//	Title       string `gorm:"not null"`
//	Description string
//	Body        string
//	Author      User
//	AuthorID    uint
//	Comments    []Comment //1:M relation, no need other decorators
//	//TODO: rename to different name, see what's affected
//	Favorites []User `gorm:"many2many:favorites;"`
//	//avoid recreation of tags: http://gorm.io/docs/associations.html
//	Tags []Tag `gorm:"many2many:article_tags;association_autocreate:false"`
//}
//
//type Comment struct {
//	gorm.Model
//	Article   Article
//	ArticleID uint
//	User      User
//	UserID    uint
//	Body      string
//}
//
//type Tag struct {
//	gorm.Model
//	Tag      string    `gorm:"unique_index"`
//	Articles []Article `gorm:"many2many:article_tags;"`
//}
