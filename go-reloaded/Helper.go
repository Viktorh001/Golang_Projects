package main

import (
	"strconv"
	"strings"
)

func Helper(s string) (string, int) {

	s = strings.TrimSpace(s)

	s = strings.TrimPrefix(s, "(")
	s = strings.TrimSuffix(s, ")")

	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return "", 0
	}

	tag := strings.TrimSpace(parts[0])
	numStr := strings.TrimSpace(parts[1])

	n, err := strconv.Atoi(numStr)
	if err != nil {
		return tag, 0
	}
	return tag, n

}
