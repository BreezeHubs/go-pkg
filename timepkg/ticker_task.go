package timepkg

import (
	"context"
	"time"
)

func TickerTaskV1(ctx context.Context, taskFunc func(ctx context.Context) error, t time.Duration) (err error) {
	ticker := time.NewTicker(t)
	defer ticker.Stop()

	if taskErr := taskFunc(ctx); taskErr != nil {
		err = taskErr
		return
	}

	done := false
	for !done {
		select {
		case <-ctx.Done():
			done = true
			err = ctx.Err()
		case <-ticker.C:
			if taskErr := taskFunc(ctx); taskErr != nil {
				done = true
				err = taskErr
			}
		}
	}
	return
}

func TickerTaskWithThrowError(ctx context.Context, taskFunc func(ctx context.Context) error, t time.Duration) error {
	if err := taskFunc(ctx); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(t):
			if err := taskFunc(ctx); err != nil {
				return err
			}
		}
	}
}

func TickerTaskWithPrintError(ctx context.Context, taskFunc func(ctx context.Context) error, errFunc func(err error), t time.Duration) error {
	errFunc(taskFunc(ctx))
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(t):
			errFunc(taskFunc(ctx))
		}
	}
}
