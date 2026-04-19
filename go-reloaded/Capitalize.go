package main

import "strings"

func ToCapitalCase(s string) string {
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
