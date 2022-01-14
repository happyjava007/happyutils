package date_utils

import (
	"testing"
	"time"
)

var ms int64 = 1642162545788
var ns int64 = 1642162545788333333
var sec int64 = 1642162545

func TestCurrTimestampMs(t *testing.T) {
	t.Log(CurrTimestampMs())
}

func TestCurrTimestampNs(t *testing.T) {
	t.Log(CurrTimestampNs())
}

func TestCurrTimestampSec(t *testing.T) {
	t.Log(CurrTimestampSec())
}

func TestMsToTime(t *testing.T) {
	timestampMs := 1531293019234
	t.Log(MsToTime(int64(timestampMs)))
}

func TestNsToTime(t *testing.T) {
	var ns int64 = 1642159791724240200
	tm := NsToTime(ns)
	t.Log(tm)
}

func TestTimeToDateTimeStr(t *testing.T) {
	t.Log(TimeToDateTimeStr(time.Now()))
}

func TestTimeToDateStr(t *testing.T) {
	t.Log(TimeToDateStr(time.Now()))
}

func TestMsToDateTimeStr(t *testing.T) {
	t.Log(MsToDateTimeStr(ms))
}

func TestMsToDateStr(t *testing.T) {
	t.Log(MsToDateStr(ms))
}

func TestNsToDateTimeStr(t *testing.T) {
	t.Log(NsToDateTimeStr(ns))
}

func TestNsToDateStr(t *testing.T) {
	t.Log(NsToDateStr(ns))
}

func TestSecToTime(t *testing.T) {
	t.Log(SecToTime(sec))
}

func TestSecToDateTimeStr(t *testing.T) {
	t.Log(SecToDateTimeStr(sec))
}

func TestSecToDateStr(t *testing.T) {
	t.Log(SecToDateStr(sec))
}