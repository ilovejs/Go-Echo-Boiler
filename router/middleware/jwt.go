package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"onsite/utils"
)

type (
	JWTConfig struct {
		Skipper    Skipper
		SigningKey interface{}
	}
	Skipper      func(c echo.Context) bool
	jwtExtractor func(echo.Context) (string, error)
)

var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusUnauthorized, "missing or malformed jwt")
	ErrJWTInvalid = echo.NewHTTPError(http.StatusForbidden, "invalid or expired jwt")
)

// todo: check MiddlewareFunc documentation
func JWT(key interface{}) echo.MiddlewareFunc {
	// see type definition above
	c := JWTConfig{}
	c.SigningKey = key
	// pass on to actual definition below
	return jwtWithConfig(c)
}

func jwtWithConfig(config JWTConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			extractor := jwtFromHeader("Authorization", "Token")
			// todo: nice go pattern
			auth, err := extractor(c)
			// no token extracted
			if err != nil {
				// use skipper
				if config.Skipper != nil {
					if config.Skipper(c) {
						return next(c)
					}
				}
				return c.JSON(http.StatusUnauthorized, utils.NewError(err))
			}
			// parse auth header
			token, parseErr := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
				// todo: nice code
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				// signed key
				return config.SigningKey, nil
			})
			fmt.Println("Jwt middleware extract: ", token)
			if parseErr != nil {
				return c.JSON(http.StatusForbidden, utils.NewError(ErrJWTInvalid))
			}

			// todo: nice go shortcut
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// read id from claims
				userID := int(claims["id"].(float64))
				// set user to context is not common
				c.Set("user", userID)
				fmt.Println("JWT: Set user to context")
				return next(c)
			}
			return c.JSON(http.StatusForbidden, utils.NewError(ErrJWTInvalid))
		}
	}
}

// jwtFromHeader returns a `jwtExtractor` that extracts token from the request header.
func jwtFromHeader(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", ErrJWTMissing
	}
}
