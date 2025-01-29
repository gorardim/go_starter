package utils

import "time"

const (
	DateTimeFmt = "2006-01-02 15:04:05"
	DateFmt     = "2006-01-02"
	MongoFmt    = "2006/01/02 15:00"
)

func StartTime(t time.Time) time.Time {
	t, _ = time.Parse(DateTimeFmt, t.Format(DateFmt)+" 00:00:00")
	return t
}

func EndTime(t time.Time) time.Time {
	t, _ = time.Parse(DateTimeFmt, t.Format(DateFmt)+" 23:59:59")
	return t
}

func DiffDays(start, end time.Time) int {
	return int(StartTime(end).Sub(StartTime(start)).Hours() / 24)
}
