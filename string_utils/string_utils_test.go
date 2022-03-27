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
	substring, err := Substring(str, 2, 5)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(substring)
}

func TestSubstringBetween(t *testing.T) {
	str := "哈Happyjava皮牛逼啊"
	between, err := SubstringBetween(str, "Ha", "牛逼")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(between)

	str = "https://www.baidu.com/apple"
	substringBetween, err := SubstringBetween(str, "//", "/")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(substringBetween)
}

//
func TestSubStringBetweenExclude(t *testing.T) {
	str := "哈Happyjava皮牛逼啊"
	exclude, _ := SubstringBetweenExclude(str, "Ha", "牛逼")
	t.Log(exclude)

	str = "https://www.baidu.com/apple"
	exclude, _ = SubstringBetweenExclude(str, "//", "/")
	t.Log(exclude)

	t.Log(SubstringBetweenExclude("哈皮世界第一啊", "第一", "哈皮"))

	t.Log(SubstringBetweenExclude("哈皮世界", "哈皮", "哈皮"))

}
