package main

import (
	"strconv"
	"strings"
)

func Processor(s string) string {
	words := strings.Fields(s)
	result := []string{}

	for i := 0; i < len(words); i++ {
		word := words[i]

		if i+1 < len(words) && (strings.HasPrefix(word, "(up,") || strings.HasPrefix(word, "(low,") || strings.HasPrefix(word, "(cap,")) {
			fulltag := word + " " + words[i+1]
			tag, n := Helper(fulltag)

			if n > len(result) {
				n = len(result)
			}

			for j := len(result) - n; j < len(result); j++ {

				switch tag {
				case "up":
					result[j] = ToUpperCase(result[j])

				case "low":
					result[j] = ToLowerCase(result[j])

				case "cap":
					result[j] = ToCapitalCase(result[j])
				}
			}
			i++
			continue
		}

		switch word {
		case "(hex)":
			if len(result) > 0 {
				val, err := HexToDec(result[len(result)-1])
				if err == nil {
					result[len(result)-1] = strconv.FormatInt(val, 10)
				}
			}
		case "(bin)":
			if len(result) > 0 {
				val, err := BintoDec(result[len(result)-1])
				if err == nil {
					result[len(result)-1] = strconv.FormatInt(val, 10)
				}
			}

		case "(low)":
			if len(result) > 0 {
				result[len(result)-1] = ToLowerCase(result[len(result)-1])
			}

		case "(up)":
			if len(result) > 0 {
				result[len(result)-1] = ToUpperCase(result[len(result)-1])
			}

		case "(cap)":
			if len(result) > 0 {
				result[len(result)-1] = ToCapitalCase(result[len(result)-1])
			}
		default:
			result = append(result, word)
		}
	}
	return FormatText(strings.Join(result, " "))
}
