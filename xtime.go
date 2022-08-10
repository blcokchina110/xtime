package xtime

import "time"

//between end-start return milliseconds
func Between(start, end time.Time) int64 {
	return end.Sub(start).Milliseconds()
}

//milli second
func MilliSecond() int64 {
	return time.Now().UnixNano() / 1e6
}

//second
func Second() int64 {
	return time.Now().Unix()
}
