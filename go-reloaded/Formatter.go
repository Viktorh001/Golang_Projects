package main

import (
	"regexp"
)

func fixSpacing(s string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(s, " ")
}

func fixPunctuation(s string) string {
	re := regexp.MustCompile(`([.,!?:;])\s+([.,!?;:])`)
	for re.MatchString(s) {
		s = re.ReplaceAllString(s, "$1$2")
	}

	re = regexp.MustCompile(`([.,!?;:'])\s+`)
	s = re.ReplaceAllString(s, "$1")

	re = regexp.MustCompile(`\s+([.,!?;:'])`)
	s = re.ReplaceAllString(s, "$1")

	re = regexp.MustCompile(`([.,!?;:]+)([^\s.,!?;:])`)
	s = re.ReplaceAllString(s, "$1 $2")

	return s
}

func fixQuotes(s string) string {
	re := regexp.MustCompile(`"\s*(.*?)\s*"`)
	s = re.ReplaceAllString(s, `"$1"`)

	re = regexp.MustCompile(`'\s*(.*?)\s*'`)
	s = re.ReplaceAllString(s, `' $1'`)
	return s

}

func fixArticles(s string) string {
	re := regexp.MustCompile(`\b([Aa])\s+([aeiouhAEIOUH])`)
	return re.ReplaceAllString(s, "${1}n $2")
}

func FormatText(s string) string {
	s = fixQuotes(s)
	s = fixArticles(s)
	s = fixPunctuation(s)
	s = fixSpacing(s)

	return s
}
