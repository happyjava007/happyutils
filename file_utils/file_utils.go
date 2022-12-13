package file_utils

import (
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func Mkdir(path string) error {
	if Exists(path) {
		return nil
	}
	return os.MkdirAll(path, os.ModePerm)
}

func Rename(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}
