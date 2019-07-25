package handler

import (
	"github.com/labstack/echo"
	. "onsite/glue"
	"onsite/router/middleware"
	"onsite/utils"
)

// self designed struct, but better to be dynamic / anonymous struct ?
type Handler struct {
	projectStore    ProjectStoreInterface
	userStore       UserStoreInterface
	basicTradeStore BasicTradeStoreInterface
	claimStore      ClaimStoreInterface
}

// DI, inject store dependencies to controller
func NewHandler(ps ProjectStoreInterface, us UserStoreInterface, bts BasicTradeStoreInterface,
	cs ClaimStoreInterface) *Handler {
	return &Handler{
		projectStore:    ps,
		userStore:       us,
		basicTradeStore: bts,
		claimStore:      cs,
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

	basicTrades := v1.Group("/bts")
	basicTrades.POST("", h.CreateBasicTrade)
	basicTrades.GET("", h.ListBasicTrades)
	basicTrades.GET("/:id", h.ReadBasicTrade)
	basicTrades.PUT("/:id", h.UpdateBasicTrade)
	basicTrades.DELETE("/:id", h.DeleteBasicTrade)

}
