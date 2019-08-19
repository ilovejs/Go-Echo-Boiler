package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

// custom new router, used by main.go
func New() *echo.Echo {

	e := echo.New()

	// for vue app: https://github.com/scotch-io/go-echo-vue-single-page-app/blob/master/todo.go
	e.File("/", "public/index.html")
	// https://echo.labstack.com/cookbook/google-app-engine#app-yaml-configuration-file
	// e.Static("/", "public")

	// log level
	e.Logger.SetLevel(log.DEBUG)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"Access-Token", // preflight Options in FE inclu such header
			// echo.HeaderXRequestedWith,
		},
		AllowMethods: []string{
			echo.GET, echo.HEAD, echo.PUT, echo.OPTIONS,
			echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	// validator v9
	e.Validator = NewValidator()
	return e
}
