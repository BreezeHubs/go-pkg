package etcdsvcdispkg

import (
	"math/rand"
	"sync/atomic"
)

// RandLB 随机算法
func RandLB(i int) int {
	return rand.Intn(i) // [0, n)
}

var (
	roundRobinLBIndex int64 = -1
	roundRobinLBLen   int64
)

// RoundRobinLB 轮巡算法
func RoundRobinLB(i int64) int64 {
	if i <= 1 {
		return 0
	}

	if roundRobinLBLen != i {
		atomic.StoreInt64(&roundRobinLBLen, i)
		atomic.StoreInt64(&roundRobinLBIndex, -1)
	} else if roundRobinLBIndex > i*2 {
		atomic.StoreInt64(&roundRobinLBIndex, -1)
	}

	atomic.AddInt64(&roundRobinLBIndex, 1)
	return roundRobinLBIndex % i
}
