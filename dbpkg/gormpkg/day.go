package gormpkg

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Day format json time field by myself
type Day struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t Day) MarshalJSON() ([]byte, error) {
	if (t == Day{}) {
		formatted := fmt.Sprintf("\"%s\"", "")
		return []byte(formatted), nil
	} else {
		formatted := fmt.Sprintf("\"%s\"", t.Format(time.DateOnly))
		return []byte(formatted), nil
	}
}

// Value insert timestamp into mysql need this function.
func (t Day) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *Day) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Day{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
