package main

import (
	"strconv"
	"strings"
)

func Processor(s string) string {
	words := strings.Fields(s)
	result := []string{}

	//loop through
	for i := 0; i < len(words); i++ {
		switch words[i] {
		//Hextodecimal
		case "(hex)":
			if len(result) > 0 {
				value, err := HexToDec(result[len(result)-1])
				if err == nil {
					result[len(result)-1] = strconv.FormatInt(value, 10)
				}
			}
		case "(bin)":
			if len(result) > 0 {
				value, err := BintoDec(result[len(result)-1])
				if err == nil {
					result[len(result)-1] = strconv.FormatInt(value, 10)
				}
			}
		default:
			result = append(result, words[i])
		}
	}
	return strings.Join(result, " ")
}
