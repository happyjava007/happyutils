package random_utils

import (
	"github.com/go-basic/uuid"
	"strings"
)

// eg:8958a8a7901d4d52bbc6c6ff7c8d4926
func UUID() string {
	return strings.ReplaceAll(uuid.New(), "-", "")
}

// eg:8958a8a7-901d-4d52-bbc6-c6ff7c8d4926
func UUIDWithRail() string {
	return uuid.New()
}
