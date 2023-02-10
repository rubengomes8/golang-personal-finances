package utils

import "time"

func DateStringToTime(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func TimeToStringDate(t time.Time) string {
	return t.Format("2006-01-02")
}
