package glue

import "onsite/model"

//type ArticleStoreInterface interface {
//	GetBySlug(string) (*model.Article, error)
//	GetUserArticleBySlug(userID uint, slug string) (*model.Article, error)
//	CreateArticle(*model.Article) error
//	UpdateArticle(*model.Article, []string) error
//	DeleteArticle(*model.Article) error
//	List(offset, limit int) ([]model.Article, int, error)
//	ListByTag(tag string, offset, limit int) ([]model.Article, int, error)
//	ListByAuthor(username string, offset, limit int) ([]model.Article, int, error)
//	ListByWhoFavorited(username string, offset, limit int) ([]model.Article, int, error)
//	ListFeed(userID uint, offset, limit int) ([]model.Article, int, error)
//	AddComment(*model.Article, *model.Comment) error
//	GetCommentsBySlug(string) ([]model.Comment, error)
//	GetCommentByID(uint) (*model.Comment, error)
//	DeleteComment(*model.Comment) error
//	AddFavorite(*model.Article, uint) error
//	RemoveFavorite(*model.Article, uint) error
//	ListTags() ([]model.Tag, error)
//}

type UserStoreInterface interface {
	GetByID(uint) (*model.User, error)
	GetByEmail(string) (*model.User, error)
	GetByUsername(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
	AddFollower(user *model.User, followerID uint) error
	RemoveFollower(user *model.User, followerID uint) error
	IsFollower(userID, followerID uint) (bool, error)
}

type ProjectStoreInterface interface {
	//CRUD
	CreateProject(*model.Project) error
	ReadProject(uint) (*model.Project, error)
	UpdateProject(uint) (*model.Project, error) //EditProject
	DeleteProject(uint, string) (*model.Project, error)
	//
	GetContractors(uint)
	GetTradeSummary(uint)
	GetTradeItems(uint, uint)            //assignTradeList(tid,pid)
	GetTradeDetailByTradeId(uint, uint)  //GET addTradeItem
	AddTradeItem(trade *model.TradeItem) //POST addTradeItem
	GetTradeItemsByProjectId(uint)       //getSavedTradeItems(pid, tid)
	//todo:TradeItemsPartialSend
	DeleteTradeItem(uint, string) //(id, reason)
	UpdateTradeItems()            //EditTradeItem 2 methods
	//Claim:
	ProjectClaim(uint)              //Merge ctr with other ?
	UpdateClaim(claim *model.Claim) //POST ProjectClaim
	UpdateClaimByID(uint, uint)     //TradeItemsClaimEdit(pid, tid)
	TradeTotalByProjectId(uint)     //GetSavedTradeTotal
	TradeItemsClaimSummary(uint)    //getClaimSummaryByMonth
	GetClaimByProjectId(uint)       //TradeItemsClaim
	GetTradeDetail(uint)            //pid
	TradeItemsClaimList(uint, uint) //GetTradeItems(uint, uint)
	GetContractorClaimReport(uint)
	SubmitClaim(uint, string, string)      //pid, month, claim number
	GetClaimHistory(uint)                  //claimId
	GetTradeItemListByLevel(uint, string)  //pid, level
	GetProjectTradeItemByLevel(uint, uint) //pid, level
	GetClaimReport(uint)
}
