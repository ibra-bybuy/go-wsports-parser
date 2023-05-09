package datetime

import "time"

// 2006-01-02
func YearMonthDay(t time.Time) string {
	return t.Format("2006-01-02")
}

// 2006-01-02T14:00:00.000Z
func Full(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000Z")
}

// 2006-01-02T14:00:00.000Z
func FromFull(t string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05.000Z", t)
}
