package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Role struct {
	gorm.Model
	Type  string
	Users []*User `gorm:"many2many:user_roles;"`
}

type User struct {
	gorm.Model
	UserName    *string `gorm:"unique;not null"`
	FirstName   string
	LastName    string
	Company     string
	Email       string `gorm:"unique;not null"`
	Mobile      string
	Password    string
	Token       string
	TokenExpire time.Time
	IsActive    bool
	IsDeleted   bool
	Bio         string
	Avatar      string
	Roles       []Role `gorm:"many2many:user_roles;"`

	//1 User has many projects
	ManageProjects      []Project       `gorm:"foreignkey:Manager"`
	CreatedProjects     []Project       `gorm:"foreignkey:CreatedBy"`
	CreatedTradeItems   []TradeItemType `gorm:"foreignkey:CreatedBy"`
	CreatedClaimHistory []ClaimHistory  `gorm:"foreignkey:CreatedBy"`

	ContractorProjects []*ContractorProject
}

type Project struct {
	gorm.Model
	ManagerID            int //fk
	Name                 string
	GrossItemBreakDown   float32
	GrossContractorClaim float32
	ContractValue        float32
	IsDeleted            bool
	SerialNo             string
	Detail               string
	QuantitySurveyor     string
	Note                 string
	TradeItems           []TradeItem //?

	PrimaryContactID int
	PrimaryContact   User

	SecondaryContactID int
	SecondaryContact   User

	ContractorProjects []*ContractorProject

	CreatedBy int //fk
}

//linking table with extra fields, this is native table with no gorm.model
type ContractorProject struct {
	ID int

	User   *User //not to create a new one
	UserID int

	Project   *Project
	ProjectID int

	SideNote  string
	IsDeleted bool

	Created time.Time
	Updated time.Time
}

//table name
func (*ContractorProject) TableName() string {
	return "contractor_project"
}

// 1 type: M tradeItem
type TradeItemType struct {
	gorm.Model
	Name      string
	IsDeleted bool
	//won't create table
	TradeItems []TradeItem
}

type TradeItem struct {
	gorm.Model
	TradeItemTypeID   int //fk
	ProjectID         int //fk
	Level             int
	DescriptionOfWork string
	ItemBreakDown     float32
	IsDeleted         bool
	LastUpdate        time.Time
	TempCheck         bool
	CreatedBy         int //fk
}

//1 item: M claim
type Claim struct {
	gorm.Model
	TradeItemID     int //fk
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
	ClaimID         int
	PreviousClaimed float32
	CreatedBy       int
}
