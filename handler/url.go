package handler

import (
	"github.com/labstack/echo"
	"onsite/router/middleware"
	"onsite/store"
	"onsite/utils"
)

// self designed struct, but better to be dynamic / anonymous struct ?
type Handler struct {
	projectStore       store.ProjectStoreInterface
	userStore          store.UserStoreInterface
	TradeCategoryStore store.TradeCategoryStoreInterface
	TradeStore         store.TradeStoreInterface
	claimStore         store.ClaimStoreInterface
}

// DI, inject store dependencies to controller
func NewHandler(
	ps store.ProjectStoreInterface,
	us store.UserStoreInterface,
	bts store.TradeCategoryStoreInterface,
	ts store.TradeStoreInterface,
	cs store.ClaimStoreInterface) *Handler {

	return &Handler{
		projectStore:       ps,
		userStore:          us,
		TradeCategoryStore: bts,
		TradeStore:         ts,
		claimStore:         cs,
	}
}

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)

	users := v1.Group("/auth")
	users.POST("/register", h.SignUp)
	users.POST("/login", h.Login)
	users.POST("/logout", h.Logout)

	// single user with api protected
	user := v1.Group("/user", jwtMiddleware)
	user.GET("/info", h.CurrentUser)
	user.PUT("", h.UpdateUser)
	user.GET("", h.ListUser)

	profiles := v1.Group("/profiles", jwtMiddleware)
	profiles.GET("/:username", h.GetProfile)
	// todo: profile

	// todo: roles

	projects := v1.Group("/projects", jwtMiddleware)
	projects.POST("", h.CreateProject)
	projects.GET("", h.ListProject)
	projects.GET("/:id", h.ReadProject)
	projects.PUT("/:id", h.UpdateProject)
	projects.DELETE("/:id", h.DeleteProject)
	// projects.POST("/:id/addtrade", h.AddTrades)

	tradeCategories := v1.Group("/bts", jwtMiddleware)
	tradeCategories.POST("", h.CreateTradeCategory)
	tradeCategories.GET("", h.ListTradeCategories)
	tradeCategories.GET("/:id", h.ReadTradeCategory)
	tradeCategories.PUT("/:id", h.UpdateTradeCategory)
	tradeCategories.DELETE("/:id", h.DeleteTradeCategory)

	trade := v1.Group("/trades", jwtMiddleware)
	trade.POST("", h.CreateTrade)
	trade.GET("", h.ListTrade)
	trade.GET("/:id", h.ReadTrade)
	trade.PUT("/:id", h.UpdateTrade)
	trade.DELETE("/:id", h.DeleteTrade)

}
