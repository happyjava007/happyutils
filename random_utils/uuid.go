package random_utils

import (
	"github.com/go-basic/uuid"
)

func UUID() string {
	return uuid.New()
}
