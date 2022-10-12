package utils

import "strings"

func StrInArray(str string, arr []string) bool {
	for _, v := range arr {
		return strings.EqualFold(str, v)
	}

	return false
}
