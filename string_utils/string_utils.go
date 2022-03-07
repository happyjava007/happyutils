package string_utils

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func IsBlank(s string) bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range s {
		if !unicode.IsSpace(v) {
			return false
		}
	}
	return true
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

func Length(s string) int {
	return utf8.RuneCountInString(s)
}

func EqualsAny(s1 string, s2 ...string) bool {
	for _, v := range s2 {
		if s1 == v {
			return true
		}
	}
	return false
}