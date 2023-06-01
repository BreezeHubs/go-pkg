package svcdispkg

// IRegister 服务注册通用接口
type IRegister interface {
	// Register 注册服务
	Register(service IService) error

	// DeRegister 注销服务
	DeRegister() error

	Close() error

	HasLease() bool
}
