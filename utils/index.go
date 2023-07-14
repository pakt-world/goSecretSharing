package utils

import (
	"log"
	"strconv"
)

func StringToInt(s string) int {
	st, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return int(st)
}
