package svcdispkg

type IDiscovery interface {
	// GetServiceAddr 获取服务的一个实例地址
	GetServiceAddr(serviceName string, f func(i int64) int64) (string, error)

	// WatchService 监控服务的地址变化
	WatchService(serviceName string) error
}
