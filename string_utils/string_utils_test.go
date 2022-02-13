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

func TestSubString(t *testing.T) {
	str := "哈皮牛逼啊"
	subString := SubString(str, 2, 5)
	t.Log(subString)
}

func TestSubStringBetween(t *testing.T) {
	str := "哈Happyjava皮牛逼啊"
	between := SubStringBetween(str, "Ha", "牛逼")
	t.Log(between)

	str = "https://www.baidu.com/apple"
	between = SubStringBetween(str, "//", "/")
	t.Log(between)
}

func TestSubStringBetweenExclude(t *testing.T) {
	str := "哈Happyjava皮牛逼啊"
	exclude := SubStringBetweenExclude(str, "Ha", "牛逼")
	t.Log(exclude)

	str = "https://www.baidu.com/apple"
	exclude = SubStringBetweenExclude(str, "//", "/")
	t.Log(exclude)
}
