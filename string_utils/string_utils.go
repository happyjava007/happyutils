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

// SubString [start,end)
func SubString(str string, start, end int) string {
	length := Length(str)
	if start < 0 || end > length || start > end {
		return ""
	}
	if start == 0 && end == length {
		return str
	}
	runes := []rune(str)
	res := runes[start:end]
	return string(res)
}

// SubStringBetween include s1 and s2
func SubStringBetween(str, s1, s2 string) string {
	return s1 + SubStringBetweenExclude(str, s1, s2) + s2
}

func SubStringBetweenExclude(str, s1, s2 string) string {
	i1 := strings.Index(str, s1)
	if i1 < 0 {
		return ""
	}
	next := str[i1+len(s1):]
	i2 := strings.Index(next, s2)
	if i2 < 0 {
		return ""
	}
	return next[:i2]
}
