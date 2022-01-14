package date_utils

import (
	"time"
)

func CurrTimestampMs() int64 {
	return TimeToMs(time.Now())
}

func CurrTimestampNs() int64 {
	return TimeToNs(time.Now())
}

func CurrTimestampSec() int64 {
	return TimeToSec(time.Now())
}

func TimeToMs(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

func TimeToNs(t time.Time) int64 {
	return t.UnixNano()
}

func TimeToSec(t time.Time) int64 {
	return t.Unix()
}

func MsToTime(timestampMs int64) time.Time {
	ns := (timestampMs % 1000) * 1000000
	return time.Unix(timestampMs/1000, ns)
}

func NsToTime(timestampNs int64) time.Time {
	sec := timestampNs / 1000000000
	ns := timestampNs % 1000000000
	return time.Unix(sec, ns)
}

func SecToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

func TimeToDateTimeStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func TimeToDateStr(t time.Time) string {
	return t.Format("2006-01-02")
}

func MsToDateTimeStr(ms int64) string {
	return TimeToDateTimeStr(MsToTime(ms))
}

func MsToDateStr(ms int64) string {
	return TimeToDateStr(MsToTime(ms))
}

func NsToDateTimeStr(ns int64) string {
	return TimeToDateTimeStr(NsToTime(ns))
}

func NsToDateStr(ns int64) string {
	return TimeToDateStr(NsToTime(ns))
}

func SecToDateTimeStr(sec int64) string {
	return TimeToDateTimeStr(SecToTime(sec))
}

func SecToDateStr(sec int64) string {
	return TimeToDateStr(SecToTime(sec))
}
