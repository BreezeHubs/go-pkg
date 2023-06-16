package timepkg

import (
	"errors"
	"time"
)

func TimestampToDatetime(timestamp int64, layout ...string) string {
	defaultlayout := time.DateTime
	if len(layout) > 0 {
		defaultlayout = layout[0]
	}

	return time.Unix(timestamp, 0).Format(defaultlayout)
}

func DateTimeToTimestamp(datetime string, layout ...string) (int64, error) {
	defaultlayout := time.DateTime
	if len(layout) > 0 {
		defaultlayout = layout[0]
	}

	timestamp, err := time.ParseInLocation(defaultlayout, datetime, time.Local)
	if err != nil {
		return 0, errors.New("请输入正确格式的时间")
	}
	return timestamp.Unix(), nil
}

func IsTodayTimestamp(timestamp int64) bool {
	now := time.Now()
	t := time.Unix(timestamp, 0).In(now.Location())
	return t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day()
}
