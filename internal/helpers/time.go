package helpers

import "time"

func NowTime() time.Time {
	return time.Now()
}

func ParseTime(value string) (time.Time, error) {
	return time.Parse(time.RFC3339, value)
}
