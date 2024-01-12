package helper

import "strconv"

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func MustInt(s string) int {
	i, err := strconv.Atoi(s)
	CheckErr(err)
	return i
}
