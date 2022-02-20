package timeHelper

import "time"

func Now() time.Time {
	return time.Now()
}

func Parse(value string) (time.Time, error) {
	return time.Parse(time.RFC3339, value)
}
