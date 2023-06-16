package timepkg

import "time"

func GetMonthListByNow(monthNum int, desc bool) []string {
	now := time.Now()
	var months []string

	if monthNum <= 0 {
		monthAgo := now.AddDate(0, monthNum, 0)
		if !desc {
			// 正序
			for date := monthAgo.AddDate(0, 1, 0); now.After(date.AddDate(0, -1, 0)); date = date.AddDate(0, 1, 0) {
				day := date.Format(YearAndMonth)
				months = append(months, day)
			}
		} else {
			// 倒序
			for date := now; monthAgo.Before(date); date = date.AddDate(0, -1, 0) {
				day := date.Format(YearAndMonth)
				months = append(months, day)
			}
		}
	} else {
		monthAfter := now.AddDate(0, monthNum, 0)
		if !desc {
			// 正序
			for date := now; monthAfter.After(date); date = date.AddDate(0, 1, 0) {
				day := date.Format(YearAndMonth)
				months = append(months, day)
			}
		} else {
			// 倒序
			for date := monthAfter.AddDate(0, -1, 0); now.Before(date.AddDate(0, 1, 0)); date = date.AddDate(0, -1, 0) {
				day := date.Format(YearAndMonth)
				months = append(months, day)
			}
		}
	}

	return months
}
