package schedule

import (
	"context"
	"testing"
)

func TestSchedule(t *testing.T) {
	schedule := NewSchedule(context.Background())
	schedule.Add("", func(ctx context.Context) error {
		for true {

		}
		return nil
	})

	err := schedule.RunAndGracefullyExit()
	if err != nil {
		t.Error(err)
	}
}
