package syspkg

import "runtime"

func SetThreadNum(n int) int {
	return runtime.GOMAXPROCS(n)
}

func GetThreadNum() int {
	return runtime.GOMAXPROCS(0)
}

// GetCPUNum 当前系统的CPU核数量
func GetCPUNum() int {
	return runtime.NumCPU()
}
