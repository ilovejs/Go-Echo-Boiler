package handler

import (
	"github.com/labstack/echo/v4"
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

	//profiles := v1.Group("/profiles", jwtMiddleware)
	//profiles.GET("/:username", h.GetProfile)
	//profiles.POST("/:username/follow", h.Follow)
	//profiles.DELETE("/:username/follow", h.Unfollow)

	//articles := v1.Group("/articles", middleware.JWTWithConfig(
	//	middleware.JWTConfig{
	//		Skipper: func(c echo.Context) bool {
	//			if c.Request().Method == "GET" && c.Path() != "/api/articles/feed" {
	//				return true
	//			}
	//			return false
	//		},
	//		SigningKey: utils.JWTSecret,
	//	},
	//))

	//articles.POST("", h.CreateArticle)
	//articles.GET("/feed", h.Feed)
	//articles.PUT("/:slug", h.UpdateArticle)
	//articles.DELETE("/:slug", h.DeleteArticle)
	//articles.POST("/:slug/comments", h.AddComment)
	//articles.DELETE("/:slug/comments/:id", h.DeleteComment)
	//articles.POST("/:slug/favorite", h.Favorite)
	//articles.DELETE("/:slug/favorite", h.Unfavorite)
	//articles.GET("", h.Articles)
	//articles.GET("/:slug", h.GetArticle)
	//articles.GET("/:slug/comments", h.GetComments)

	//tags := v1.Group("/tags")
	//tags.GET("", h.Tags)

	projects := v1.Group("/projects")
	projects.GET("/:id", h.ReadProject)
}
