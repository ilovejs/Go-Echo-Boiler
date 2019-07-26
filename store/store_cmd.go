package store

import (
	"onsite/models"
)

/* Store interfaces are injected to handler in url.go */

type UserStoreInterface interface {
	Create(*models.User) error
	GetByID(int) (*models.User, error)
	GetByEmail(string) (*models.User, error)
	GetByUserName(string) (*models.User, error)
	Update(*models.User) error
	Delete(int) error
}

type ProjectStoreInterface interface {
	Create(*models.Project) error
	Read(int) (*models.Project, error)
	Update(*models.Project) error
	Delete(int) error
	List(offset int, limit int) ([]*models.Project, int, error)
	Contractors(pid int) ([]*models.User, error)
}

type TradeCategoryStoreInterface interface {
	Create(trade *models.TradeCategory) error
	Get(ids ...int) ([]*models.TradeCategory, error) //Get one or all
	Update(trade *models.TradeCategory) error
	Delete(int) error
	List(offset int, limit int) ([]*models.TradeCategory, int, error)
}

type TradeStoreInterface interface {
	//ProjectTradeDetails(projectId int) //old: GetTradeDetail() controller
	//GetTradeSummary(int)
	//GetTradeDetailByTradeId(int, int)  //GET addTradeItem
	//GetTradeItemsByProjectId(int)       //getSavedTradeItems(pid, tid)
	//GetTradeItemListByLevel(int, string)  //pid, level
	//GetProjectTradeItemByLevel(int, int) //pid, level
	//TradeTotalByProjectId(int)     //GetSavedTradeTotal
	////todo:TradeItemsPartialSend
	//GetTradeItems(int, int)         //assignTradeList(tid,pid)
	//AddTradeItem(tradeId int, projectId int) //ProjectController
	//DeleteTradeItem(int, string)    //(id, reason)
	//UpdateTradeItems()              //EditTradeItem 2 methods
}

type ClaimStoreInterface interface {
	//TradeItemsClaimSummary(int)    //getClaimSummaryByMonth
	//ProjectClaim(int)              //Merge ctr with other ?
	//TradeClaimHistory(tradeClaimId int)  //step 4, 'view'
	////GetClaimByProjectId(int)       //TradeItemsClaim
	//UpdateClaim(claim *models.Claim) //POST ProjectClaim
	//UpdateClaimByID(int, int)     //TradeItemsClaimEdit(pid, tid)
	//TradeItemsClaimList(int, int) //GetTradeItems(int, int)
	//SubmitClaim(uint, string, string)  //pid, month, claim number
	//GetContractorClaimReport(int)
	//GetClaimReport(int)
}
