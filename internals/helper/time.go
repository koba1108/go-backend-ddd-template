package helper

import "time"

func Int64ToTime(i int64) time.Time {
	return time.Unix(i, 0)
}

func Int64ToTimePrt(i int64) *time.Time {
	t := Int64ToTime(i)
	return &t
}
