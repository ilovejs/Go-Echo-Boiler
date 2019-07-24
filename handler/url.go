package handler

import (
	"github.com/labstack/echo"
	. "onsite/glue"
	"onsite/router/middleware"
	"onsite/utils"
)

type Handler struct {
	projectStore ProjectStoreInterface
	userStore    UserStoreInterface
	claimStore   ClaimStoreInterface
}

func NewHandler(
	ps ProjectStoreInterface,
	us UserStoreInterface,
	cs ClaimStoreInterface) *Handler {
	return &Handler{
		projectStore: ps,
		userStore:    us,
		claimStore:   cs,
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

	projects := v1.Group("/projects")
	projects.POST("", h.Create)
	projects.GET("/:id", h.Read)
	//projects.PUT("/:id", h.UpdateProject)
	//projects.DELETE("/:id", h.DeleteProject)
}