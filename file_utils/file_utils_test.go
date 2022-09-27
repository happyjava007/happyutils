package file_utils

import "testing"

func TestExists(t *testing.T) {
	exists := Exists("c:/apple.txt")
	t.Log(exists)
}
