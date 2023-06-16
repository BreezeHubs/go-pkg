package timepkg

import (
	"time"
)

func GetDayListByNow(dayNum int, desc bool) []string {
	now := time.Now()
	var days []string

	if dayNum <= 0 {
		dayAgo := now.AddDate(0, 0, dayNum)
		if !desc {
			// 正序
			for date := dayAgo.AddDate(0, 0, 1); date.Before(now.AddDate(0, 0, 1)); date = date.AddDate(0, 0, 1) {
				day := date.Format(time.DateOnly)
				days = append(days, day)
			}
		} else {
			// 倒序
			for date := now; date.After(dayAgo); date = date.AddDate(0, 0, -1) {
				day := date.Format(time.DateOnly)
				days = append(days, day)
			}
		}
	} else {
		dayAfter := now.AddDate(0, 0, dayNum)
		if !desc {
			// 正序
			for date := now; dayAfter.After(date); date = date.AddDate(0, 0, 1) {
				day := date.Format(time.DateOnly)
				days = append(days, day)
			}
		} else {
			// 倒序
			for date := dayAfter.AddDate(0, 0, -1); now.Before(date.AddDate(0, 0, 1)); date = date.AddDate(0, 0, -1) {
				day := date.Format(time.DateOnly)
				days = append(days, day)
			}
		}
	}

	return days
}
