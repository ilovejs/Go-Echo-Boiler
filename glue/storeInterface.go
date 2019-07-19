package glue

type UserStoreInterface interface {
	//Create(*model.User) error
	//GetByID(uint) (*model.User, error)
	//GetByEmail(string) (*model.User, error)
	//GetByUserName(string) (*model.User, error)
	//Update(*model.User) error
	//Delete(user *model.User) error
}

type ProjectStoreInterface interface {
	//CRUD
	//CreateProject(*model.Project) (*model.Project, error)
	//ReadProject(int) (*model.Project, error)
	//UpdateProject(uint) (*model.Project, error) //EditProject
	//DeleteProject(uint, string) (*model.Project, error)

	//GetTradeSummary(uint)
	//GetTradeItems(uint, uint)            //assignTradeList(tid,pid)
	//GetTradeDetailByTradeId(uint, uint)  //GET addTradeItem
	//AddTradeItem(trade *model.TradeItem) //POST addTradeItem
	//GetTradeItemsByProjectId(uint)       //getSavedTradeItems(pid, tid)
	////todo:TradeItemsPartialSend
	//DeleteTradeItem(uint, string) //(id, reason)
	//UpdateTradeItems()            //EditTradeItem 2 methods
	////Claim:
	//ProjectClaim(uint)              //Merge ctr with other ?
	//UpdateClaim(claim *model.Claim) //POST ProjectClaim
	//UpdateClaimByID(uint, uint)     //TradeItemsClaimEdit(pid, tid)
	//TradeTotalByProjectId(uint)     //GetSavedTradeTotal
	//TradeItemsClaimSummary(uint)    //getClaimSummaryByMonth
	//GetClaimByProjectId(uint)       //TradeItemsClaim
	//GetTradeDetail(uint)            //pid
	//TradeItemsClaimList(uint, uint) //GetTradeItems(uint, uint)
	//GetContractorClaimReport(uint)
	//SubmitClaim(uint, string, string)      //pid, month, claim number
	//GetClaimHistory(uint)                  //claimId
	//GetTradeItemListByLevel(uint, string)  //pid, level
	//GetProjectTradeItemByLevel(uint, uint) //pid, level
	//GetClaimReport(uint)
}
