package schedule

import (
	"context"
	"fmt"
	"github.com/BreezeHubs/go-pkg/logpkg"
	"sync"
)

type JobsManager struct {
	mutex         sync.Mutex
	activeTargets map[string]IJob
	jobs          []IJob
}

type IJob interface {
	Hash() string
	Start() error
	Stop() error
}

func NewJobManager(jobs []IJob) *JobsManager {
	return &JobsManager{
		activeTargets: make(map[string]IJob),
		jobs:          jobs,
	}
}

func (jm *JobsManager) StopAll() error {
	jm.mutex.Lock()
	defer jm.mutex.Unlock()

	var errList []error
	for _, v := range jm.activeTargets {
		if err := v.Stop(); err != nil {
			errList = append(errList, err)
		}
	}
	if len(errList) > 0 {
		return fmt.Errorf("failed to stop all jobs: %v", errList)
	}
	return nil
}

func (jm *JobsManager) ListenAndRetry(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			if stopErr := jm.StopAll(); stopErr != nil {
				return stopErr
			}
			return ctx.Err()
		default:
			if err := jm.Sync(); err != nil {
				logpkg.ZapLogError(&logpkg.LogMessage{Tag: "pkg.JobsManager.ListenAndRedo.Sync.error", Err: err})
			}
		}
	}
}

func (jm *JobsManager) Sync() error {
	thisNewTargets := make(map[string]IJob)
	thisAllTargets := make(map[string]IJob)

	jm.mutex.Lock()
	defer jm.mutex.Unlock()

	// 将jobs中所有的IJob存入thisAllTargets
	for _, t := range jm.jobs {
		hash := t.Hash()
		thisAllTargets[hash] = t

		// 如果IJob不在activeTargets中，则说明需要新开一个IJob
		if _, ok := jm.activeTargets[hash]; !ok {
			thisNewTargets[hash] = t
			jm.activeTargets[hash] = t
		}
	}

	// 停止旧的IJob
	var errList []error
	for hash, t := range jm.activeTargets {
		if _, ok := thisAllTargets[hash]; !ok {
			if err := t.Stop(); err != nil {
				errList = append(errList, err)
			}
			delete(jm.activeTargets, hash)
		}
	}
	if len(errList) > 0 {
		return fmt.Errorf("failed to stop old jobs: %v", errList)
	}

	// 开启新的
	var wg sync.WaitGroup
	for _, t := range thisNewTargets {
		wg.Add(1)

		go func(t IJob) {
			defer wg.Done()
			if err := t.Start(); err != nil {
				logpkg.ZapLogError(&logpkg.LogMessage{Tag: "pkg.JobsManager.Sync.Start.error", Err: err})
			}
		}(t)
	}
	wg.Wait()
	return nil
}
