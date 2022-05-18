package random_utils

import "testing"

func TestRandomString(t *testing.T) {
	t.Log(RandomString(120))
	t.Log(RandomString(10))
	t.Log(RandomString(12))
	t.Log(RandomString(8))
}

func TestRandomAlphabetic(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Log(RandomAlphabetic(5))
		t.Log(RandomAlphabetic(10))
		t.Log(RandomAlphabetic(20))
	}
}

func TestRandomNumeric(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Log(RandomNumeric(5))
		t.Log(RandomNumeric(10))
		t.Log(RandomNumeric(20))
	}
}
