package utils

import "time"

func Milli2Time(milliSeconds int64) time.Time{
	return time.Unix(0,milliSeconds * int64(time.Millisecond))
}

func StartOfUTCTime(t time.Time) time.Time {
	return time.Date(t.UTC().Year(),t.UTC().Month(),t.UTC().Day(),0,0,0,0,time.UTC)
}

func Time2Milli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
