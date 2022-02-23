package helpers

import "time"

func TimeNow() time.Time {
	return time.Now()
}

func TimeParse(value string) (time.Time, error) {
	return time.Parse(time.RFC3339, value)
}
