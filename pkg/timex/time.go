package timex

import "time"

func FormatDaySpan(t time.Time) (start, end string) {
	start = t.Format("2006-01-02") + " 00:00:00"
	end = t.Format("2006-01-02") + " 23:59:59"
	return
}

func YesterdayStart() string {
	return time.Now().AddDate(0, 0, -1).Format("2006-01-02") + " 00:00:00"
}

func YesterdayEnd() string {
	return time.Now().AddDate(0, 0, -1).Format("2006-01-02") + " 23:59:59"
}

func TodayStart() string {
	return time.Now().Format("2006-01-02") + " 00:00:00"
}

func TodayEnd() string {
	return time.Now().Format("2006-01-02") + " 23:59:59"
}
