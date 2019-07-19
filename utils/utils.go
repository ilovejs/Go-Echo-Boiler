package utils

import "strings"

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
