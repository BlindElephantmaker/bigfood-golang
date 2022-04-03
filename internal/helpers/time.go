package helpers

import "time"

const (
	Day  = 24 * time.Hour
	Year = 360 * Day
)

func NowTime() time.Time {
	return time.Now()
}

// todo: not used
//func ParseTime(value string) (time.Time, error) {
//	return time.Parse(time.RFC3339, value)
//}
