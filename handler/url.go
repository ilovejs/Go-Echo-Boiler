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
	claimStore         store.ClaimStoreInterface
}

// DI, inject store dependencies to controller
func NewHandler(ps store.ProjectStoreInterface, us store.UserStoreInterface, bts store.TradeCategoryStoreInterface,
	cs store.ClaimStoreInterface) *Handler {
	return &Handler{
		projectStore:       ps,
		userStore:          us,
		TradeCategoryStore: bts,
		claimStore:         cs,
	}
}

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)

	guests := v1.Group("/users")
	guests.POST("", h.SignUp)
	guests.POST("/login", h.Login)

	user := v1.Group("/user", jwtMiddleware)
	user.GET("", h.CurrentUser)
	user.PUT("", h.UpdateUser)

	profiles := v1.Group("/profiles", jwtMiddleware)
	profiles.GET("/:username", h.GetProfile)
	//todo: profile

	projects := v1.Group("/projects")
	projects.POST("", h.CreateProject)
	projects.GET("/:id", h.ReadProject)
	projects.PUT("/:id", h.UpdateProject)
	projects.DELETE("/:id", h.DeleteProject)

	TradeCategories := v1.Group("/bts")
	TradeCategories.POST("", h.CreateTradeCategory)
	TradeCategories.GET("", h.ListTradeCategories)
	TradeCategories.GET("/:id", h.ReadTradeCategory)
	TradeCategories.PUT("/:id", h.UpdateTradeCategory)
	TradeCategories.DELETE("/:id", h.DeleteTradeCategory)

}
