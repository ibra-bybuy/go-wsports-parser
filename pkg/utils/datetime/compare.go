package datetime

import "time"

func SameDay(t1 time.Time, t2 time.Time) bool {
	if t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day() {
		return true
	}

	return false
}

func SameDayHour(t1 time.Time, t2 time.Time) bool {
	if t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day() && t1.Hour() == t2.Hour() {
		return true
	}

	return false
}

func IsTodayOrTomorrow(t time.Time) bool {
	today := SameDay(t, time.Now())
	tomorrow := SameDay(t, time.Now().Add(time.Hour*24))
	if today || tomorrow {
		return true
	}

	return false
}
