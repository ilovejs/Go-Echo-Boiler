package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	gorm.Model
	UserName    string `gorm:"unique;not null"`
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Company     string `gorm:"not null"`
	Email       string `gorm:"unique;not null"`
	Mobile      string
	Password    string `gorm:"not null"`
	Token       string
	TokenExpire time.Time
	IsActive    bool
	IsDeleted   bool
	Bio         *string
	Image       *string
	Roles       []Role `gorm:"many2many:user_roles;"`

	//1 User has many projects
	ManageProjects      []Project       `gorm:"foreignkey:Manager"`
	CreatedProjects     []Project       `gorm:"foreignkey:CreatedBy"`
	CreatedTradeItems   []TradeItemType `gorm:"foreignkey:CreatedBy"`
	CreatedClaimHistory []ClaimHistory  `gorm:"foreignkey:CreatedBy"`

	ContractorProjects []*ContractorProject
}

func (u *User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}
