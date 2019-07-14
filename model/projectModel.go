package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Role struct {
	gorm.Model
	Type string
}

type Login struct {
	gorm.Model
	UserName    string
	FirstName   string
	LastName    string
	Company     string
	Email       string
	Mobile      string
	Password    string
	Token       *string
	TokenExpire time.Time
	IsActive    bool
	IsDeleted   bool
	Bio         *string
	Avatar      *string
	RoleID      uint
	Role        []Role
	//1 login has many projects
	ManageProjects     []Project `gorm:"foreignkey:Manager"`
	CreateProjects     []Project `gorm:"foreignkey:CreatedBy"`
	ContractorProjects []*ContractorProject
}

type Project struct {
	gorm.Model
	ManagerID            uint //fk
	Name                 string
	GrossItemBreakDown   float32
	GrossContractorClaim float32
	ContractValue        float32
	IsDeleted            bool
	SerialNo             string
	Detail               string
	CreatedBy            uint //fk
	QuantitySurveyor     string
	Note                 string
	//TradeItems           []Item
	PrimaryContactID   uint
	PrimaryContact     Login
	SecondaryContactID uint
	SecondaryContact   Login
	ContractorProjects []*ContractorProject
}

type ContractorProject struct {
	ID int

	Login   Login
	LoginID uint

	Project   Project
	ProjectID uint
	//extra fields
	SideNote  string
	IsDeleted bool
	Created   *time.Time
	Updated   *time.Time
}

//table name
func (*ContractorProject) TableName() string {
	return "contractor_project"
}

// 1 type: M tradeItem
//type ItemType struct {
//	gorm.Model
//	Name      string
//	IsDeleted bool
//}
//
////Trade Item
//type Item struct {
//	gorm.Model
//	Type              ItemType `gorm:"foreignKey:ItemTypeID"`
//	Project           Project  `gorm:"foreignKey:ProjectID"`
//	Level             int
//	DescriptionOfWork string
//	ItemBreakDown     float32
//	IsDeleted         bool
//	CreatedBy         Login `gorm:"foreignKey:LoginID"`
//	LastUpdate        *time.Time
//	TempCheck         bool
//}
//
////1 item : M claim
//type Claim struct {
//	gorm.Model
//	Item            Item `gorm:"foreignKey:ItemID"`
//	GrossAmount     float32
//	ClaimedAmount   float32
//	PreviousClaimed float32
//	AmountDue       float32
//	CostToCompleted float32
//	IsDeleted       bool
//}
//
////1 claim: M history
//type ClaimHistory struct {
//	gorm.Model
//	Claim           Claim `gorm:"foreignKey:ClaimID"`
//	PreviousClaimed float32
//	CreatedBy       Login `gorm:"foreignKey:LoginID"`
//}
