package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"onsite/models"
	"strings"
)

// turn [int64 37] into 37
func UnmarshalSingleTypeValue(c string) string {
	//single arg
	c = c[1 : len(c)-1] //take out [] wraps
	s := strings.SplitN(c, " ", 2)
	//log.Printf("%T => %q\n", s, s)
	return s[1]
}

func DieIf(err error) {
	if err != nil {
		panic(err)
	}
}

func LogIf(err error) {
	if err != nil {
		log.Println(err)
	}
}

func HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func CheckPassword(u *models.User, plain string) bool {
	var pStr string
	err := u.Password.Scan(&pStr)
	if err != nil {
		log.Fatalln(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(pStr), []byte(plain))
	return err == nil
}
