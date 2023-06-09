package timepkg

import (
	"context"
	"time"
)

func TickerTask(ctx context.Context, taskFunc func(ctx context.Context) error, t time.Duration) error {
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

type tickerTaskChannel struct {
	C chan error
}

func TickerTaskWithChannel(ctx context.Context, taskFunc func(ctx context.Context) error, t time.Duration) *tickerTaskChannel {
	tc := tickerTaskChannel{
		C: make(chan error, 1),
	}

	if err := taskFunc(ctx); err != nil {
		tc.C <- err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				tc.C <- ctx.Err()
			case <-time.After(t):
				if err := taskFunc(ctx); err != nil {
					tc.C <- err
				}
			}
		}
	}()

	return &tc
}
