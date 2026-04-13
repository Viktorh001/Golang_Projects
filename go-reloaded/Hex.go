package main

import (
	"strconv"
)

func HexToDec(str string) (int64, error) {
	return strconv.ParseInt(str, 16, 64)
}
