package random_utils

import (
	"math/rand"
	"strings"
	"time"
)

var randomChar = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's',
	't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
	'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString 返回随机字符串，包含大小写字母，数字
func RandomString(length int) string {
	sb := strings.Builder{}
	for i := 0; i < length; i++ {
		n := rand.Intn(len(randomChar))
		sb.WriteByte(randomChar[n])
	}
	return sb.String()
}

// RandomAlphabetic a-z, A-Z
func RandomAlphabetic(length int) string {
	sb := strings.Builder{}
	for i := 0; i < length; i++ {
		n := rand.Intn(52)
		sb.WriteByte(randomChar[n])
	}
	return sb.String()
}

// RandomNumeric 返回随机数字 0-9
func RandomNumeric(length int) string {
	sb := strings.Builder{}
	for i := 0; i < length; i++ {
		n := rand.Intn(10)
		sb.WriteByte(randomChar[52+n])
	}
	return sb.String()
}
