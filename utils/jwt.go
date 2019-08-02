package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// this is the SECRET !!
var JWTSecret = []byte("KL:JJH&*OOH#rewrjkl jzk;f qsdar.?<:IOHQWDQW")

// used in SignUp
func GenerateJWT(uid int) string {
	// signing method
	token := jwt.New(jwt.SigningMethodHS256)
	// private content
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = uid
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()
	// sign with secret
	t, signErr := token.SignedString(JWTSecret)
	if signErr != nil {
		fmt.Println("Sign JWT token failed: ", signErr)
	}
	return t
}
