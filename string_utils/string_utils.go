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

func Substring(source string, begin, end int) (string, error) {
	if source == "" {
		return "", nil
	}
	if begin < 0 || end > Length(source) {
		return "", IndexOut
	}
	return string([]rune(source)[begin:end]), nil
}

func SubstringWithoutError(source string, begin, end int) string {
	substring, _ := Substring(source, begin, end)
	return substring
}

func SubstringBetweenExclude(source, s1, s2 string) (string, error) {
	i1 := strings.Index(source, s1)
	i2 := strings.Index(source, s2)
	if i1 < 0 {
		return "", BeginSubstringNotFound
	}
	if i2 < 0 {
		return "", EndSubstringNotFound
	}

	substring := string([]byte(source)[i1+len(s1) : len(source)])

	i2 = strings.Index(substring, s2)
	if i2 < 0 {
		return "", EndSubstringNotFound
	}

	s := string([]byte(substring)[0:i2])
	return s, nil
}

func SubstringBetweenExcludeWithoutError(source, s1, s2 string) string {
	exclude, _ := SubstringBetweenExclude(source, s1, s2)
	return exclude
}

func SubstringBetween(source, s1, s2 string) (string, error) {
	substring, err := SubstringBetweenExclude(source, s1, s2)
	if err != nil {
		return "", err
	}
	return s1 + substring + s2, nil
}

func SubstringBetweenWithoutError(source, s1, s2 string) string {
	between, _ := SubstringBetween(source, s1, s2)
	return between
}
