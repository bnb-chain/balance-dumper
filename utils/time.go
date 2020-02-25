package utils

import "time"

func Milli2Time(milliSeconds int64) time.Time{
	return time.Unix(0,milliSeconds * int64(time.Millisecond))
}

func StartOfUTCTime(t time.Time) time.Time {
	return time.Date(t.UTC().Year(),t.UTC().Month(),t.UTC().Day(),0,0,0,0,time.UTC)
}

func Time2Milli(t *time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func IsSameUTCDay(t1 time.Time,t2 time.Time) bool {
	utcT1 := t1.UTC()
	utcT2 := t2.UTC()
	return utcT1.Year() == utcT2.Year() && utcT1.Month() == utcT2.Month() && utcT1.Day() == utcT2.Day()
}
