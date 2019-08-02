package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// this is the SECRET !!
var JWTSecret = []byte("KL:JJH&*OOH#rewrjkl jzk;f qsdar.?<:IOHQWDQW")

func GenerateJWT(id int) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}
