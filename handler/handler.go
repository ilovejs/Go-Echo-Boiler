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
}

func NewHandler(ps ProjectStoreInterface, us UserStoreInterface) *Handler {
	return &Handler{
		projectStore: ps,
		userStore:    us,
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

	//projects := v1.Group("/projects")
	//projects.GET("/:id", h.ReadProject)
}
