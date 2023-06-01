package svcdispkg

// IService 服务类型
type IService interface {
	Name() string
	Addr() string
}
