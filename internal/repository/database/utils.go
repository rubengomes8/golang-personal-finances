package database

import "time"

func ToTime(unixTime int64) time.Time {
	return time.Unix(unixTime, 0)
}

func ToUnix(time time.Time) int64 {
	return time.Unix()
}
