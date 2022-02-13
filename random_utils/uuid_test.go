package random_utils

import "testing"

func TestUUID(t *testing.T) {
	t.Log(UUID())
	t.Logf(UUIDWithRail())
}
