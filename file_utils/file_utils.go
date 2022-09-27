package file_utils

import (
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func Rename(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}
