package timepkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDayListByNow(t *testing.T) {
	testCase := []struct {
		name   string
		dayNum int
		desc   bool
		want   []string
	}{
		{
			name:   "生成最近15天日期，倒序",
			dayNum: -15,
			desc:   true,
			want: []string{"2023-06-13", "2023-06-12", "2023-06-11", "2023-06-10", "2023-06-09", "2023-06-08",
				"2023-06-07", "2023-06-06", "2023-06-05", "2023-06-04", "2023-06-03", "2023-06-02", "2023-06-01",
				"2023-05-31", "2023-05-30", "2023-05-29"},
		},
		{
			name:   "生成最近2天日期，正序",
			dayNum: -2,
			desc:   false,
			want:   []string{"2023-06-11", "2023-06-12", "2023-06-13"},
		},
		{
			name:   "生成后2天日期，倒序",
			dayNum: 2,
			desc:   true,
			want:   []string{"2023-06-15", "2023-06-14", "2023-06-13"},
		},
		{
			name:   "生成后2天日期，正序",
			dayNum: 2,
			desc:   false,
			want:   []string{"2023-06-13", "2023-06-14", "2023-06-15"},
		},
	}

	for _, ts := range testCase {
		list := GetDayListByNow(ts.dayNum, ts.desc)
		assert.Equal(t, ts.want, list)
	}
}
