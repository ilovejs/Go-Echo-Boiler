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
	LoginId     uint `sql:"type:int not null"`
	UserName    string
	Email       string
	Password    string
	Token       *string
	TokenExpire time.Time
	IsActive    bool
	IsDeleted   bool
	Bio         *string
	Avatar      *string
	RoleId      uint
	Role        []Role
}

type UserProfile struct {
	gorm.Model
	Login       Login
	Mobile      string
	FirstName   string
	LastName    string
	Company     string
	Contractors []Login
}

type Project struct {
	gorm.Model
	Manager              Login
	Name                 string
	GrossItemBreakDown   float32
	GrossContractorClaim float32
	ContractValue        float32
	IsDeleted            bool
	SerialNo             string
	Detail               string
	CreatedBy            Login
	QuantitySurveyor     string
	Note                 string
	Contractors          []Login
	TradeItems           []Item
	PrimaryUserProfile   UserProfile
	SecondaryUserProfile UserProfile
}

// 1 type: M tradeItem
type ItemType struct {
	gorm.Model
	Name      string
	IsDeleted bool
}

//Trade Item
type Item struct {
	gorm.Model
	Type              ItemType
	Project           Project
	Level             int
	DescriptionOfWork string
	ItemBreakDown     float32
	IsDeleted         bool
	CreatedBy         Login
	LastUpdate        time.Time
	TempCheck         bool
}

//1 item : M claim
type Claim struct {
	gorm.Model
	Item            Item
	GrossAmount     float32
	ClaimedAmount   float32
	PreviousClaimed float32
	AmountDue       float32
	CostToCompleted float32
	IsDeleted       bool
}

//1 claim: M history
type ClaimHistory struct {
	gorm.Model
	Claim           Claim
	PreviousClaimed float32
	CreatedOn       time.Time
	CreatedBy       Login
}
