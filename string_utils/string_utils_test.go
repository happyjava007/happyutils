package string_utils

import (
	"strings"
	"testing"
)

func TestIsBlank(t *testing.T) {
	t.Log(IsBlank(""))
	t.Log(IsBlank("\t"))
	t.Log(IsBlank("\n"))
	t.Log(IsBlank("                      "))
	t.Log(IsBlank("  \t  \n"))
	t.Log(IsBlank("  apple  "))
}

func TestIsNotBlank(t *testing.T) {
	t.Log(IsNotBlank(" a \t"))
}

func TestLength(t *testing.T) {
	t.Log(len("哈皮"))
	t.Log(Length("哈皮"))
}

func TestContainsAny(t *testing.T) {
	s1 := "哈皮"
	s2 := "哈"
	t.Log(strings.ContainsAny(s1, s2))
}
