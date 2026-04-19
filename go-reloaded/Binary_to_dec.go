package main

import "strconv"

func BintoDec(s string) (int64, error) {
	return strconv.ParseInt(s, 2, 64)
}
