package store

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Role struct {
	Type string
}

type User struct {
	gorm.Model
	LoginId     uint `gorm:"primary_key" sql:"type:int not null"`
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
	UserLogin   string `gorm:"foreignkey:LoginId"`
	Mobile      string
	FirstName   string
	LastName    string
	Company     string
	Contractors []User
}

type Project struct {
	gorm.Model
	Manager              User
	Name                 string
	GrossItemBreakDown   float32
	GrossContractorClaim float32
	ContractValue        float32
	IsDeleted            bool
	SerialNo             string
	Detail               string
	CreatedBy            User
	QuantitySurveyor     string
	Note                 string
	Contractors          []User
	TradeItems           []TradeItem
	PrimaryUserProfile   UserProfile
	SecondaryUserProfile UserProfile
}

// 1 type: M tradeItem
type TradeItemType struct {
	gorm.Model
	Name      string
	IsDeleted bool
}

//TradeItem
type TradeItem struct {
	gorm.Model
	Type              TradeItemType
	Project           Project
	Level             int
	DescriptionOfWork string
	ItemBreakDown     float32
	IsDeleted         bool
	CreatedBy         User
	LastUpdate        time.Time
	TempCheck         bool
}

//1 item : M claim
type TradeItemClaim struct {
	gorm.Model
	TradeItem       TradeItem
	GrossAmount     float32
	ClaimedAmount   float32
	PreviousClaimed float32
	AmountDue       float32
	CostToCompleted float32
	IsDeleted       bool
}

type ItemClaimHistory struct {
	gorm.Model
	TradeItem      TradeItem
	TradeItemClaim TradeItemClaim
}
