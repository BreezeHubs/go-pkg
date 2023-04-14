package schedule

import (
	"context"
	"errors"
	"github.com/BreezeHubs/bekit/sys"
	"log"
	"sync"
	"time"
)

type schedule struct {
	tasks map[string]func(context.Context) error
	ctx   context.Context
	//cancel *context.CancelFunc
}

func NewSchedule(ctx context.Context) *schedule {
	s := &schedule{
		tasks: make(map[string]func(context.Context) error, 8),
	}
	ctx, _ = context.WithCancel(ctx)
	s.ctx = ctx
	//s.cancel = &cancel
	return s
}

func (s *schedule) Add(name string, f func(context.Context) error) {
	s.tasks[name] = f
}

func (s *schedule) RunAndGracefullyExit() error {
	errChan := make(chan error, 1)

	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	//创建退出信号监听
	signal := sys.NewListenExitSignal()

	//监听退出信号
	go func() {
		for !signal.IsExit() {
			time.Sleep(10 * time.Millisecond)
		}
		errChan <- errors.New("listen exit signal: out")
		cancel()
	}()

	// 编排开始
	var wg sync.WaitGroup
	for name, task := range s.tasks {
		wg.Add(1)
		name := name
		task := task

		go func() {
			defer wg.Done()

			errC := make(chan error, 1)
			go func() {
				if err := task(ctx); err != nil {
					errC <- err
				}
			}()

			select {
			case err := <-errC:
				cancel()
				errChan <- errors.New(name + " " + err.Error())
			case <-ctx.Done():
			}
		}()
		log.Println("Add task", name)
	}
	wg.Wait()
	return <-errChan
}
