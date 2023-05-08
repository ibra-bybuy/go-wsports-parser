package datetime

import (
	"strconv"
	"time"
)

func UnixToTime(t int64) time.Time {
	tm := time.Unix(t, 0)
	return tm
}

func StrUnixToTime(str string) (time.Time, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return time.Now(), err
	}

	return UnixToTime(i), nil
}
