package file_utils

import "testing"

func TestExists(t *testing.T) {
	exists := Exists("c:/apple.txt")
	t.Log(exists)
}

func TestIsFile(t *testing.T) {
	isFile := IsFile("c:/apple.txt")
	t.Log(isFile)
}
