package datetime

import (
	"strconv"
	"time"
)

func UnixTimeToDate(sec int64, format string) string {
	t := time.Unix(sec, 0)
	return t.Format(format)
}

func UnixTimeStringToTime(unixtime string) time.Time {
	unixsec, err := strconv.ParseInt(unixtime, 10, 64)
	if err != nil {
		return time.Now()
	}
	return time.Unix(unixsec, 0)
}
