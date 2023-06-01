package etcdsvcdispkg

import (
	"context"
	"github.com/BreezeHubs/go-pkg/svcdispkg"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
	"time"
)

type SvcDisClient struct {
	cli *clientv3.Client // 客户端信息

	leaseID clientv3.LeaseID // 租约信息，基于租约来做健康监测

	leaseTTL time.Duration // 健康监测间隔时间，单位s

	ctx context.Context

	lock sync.Mutex
}

const leaseTTL = 5

func NewSvcDisClient(ctx context.Context, conf *EtcdSvcDisConfig) (svcdispkg.IRegister, error) {
	if conf == nil || len(conf.Endpoints) == 0 {
		return nil, errors.New("endpoints is nil")
	}

	/// 连接 etcd
	//clientv3.Config{
	//	Endpoints:            endpoints,
	//	Context:              ctx,
	//	DialTimeout:          5 * time.Second,
	//	DialKeepAliveTime:    10 * time.Second,
	//	DialKeepAliveTimeout: 5 * time.Second,
	//	//Username:             "",
	//	//Password:             "",
	//}
	cli, err := clientv3.New(*conf)
	if err != nil {
		return nil, err
	}

	c, _ := context.WithTimeout(conf.Context, time.Second)
	if _, err = cli.Put(c, "svcDisClient-test-key", "1"); err != nil {
		return nil, err
	}

	return &SvcDisClient{
		cli:      cli,
		leaseTTL: leaseTTL,
		ctx:      ctx,
	}, nil
}

// Register 注册服务
func (r *SvcDisClient) Register(service svcdispkg.IService) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	// 租约申请
	resp, err := r.cli.Grant(r.ctx, int64(r.leaseTTL.Seconds()))
	if err != nil {
		return errors.Wrap(err, "租约申请失败")
	}
	r.leaseID = resp.ID // 租约ID

	// 将服务地址 put 到 etcd，同时绑定租约
	serviceName := service.Name() + "-" + uuid.New().String()
	_, err = r.cli.Put(r.ctx, serviceName, service.Addr(), clientv3.WithLease(r.leaseID))
	if err != nil {
		return errors.Wrap(err, "租约绑定失败")
	}

	// 租期续约
	go r.keepAlive()
	return nil
}

// 维持租约
func (r *SvcDisClient) keepAlive() {
	c, err := r.cli.KeepAlive(r.ctx, r.leaseID)
	if err != nil {
		errors.Wrap(err, "启动维持租约失败")
	}

	for {
		if _, ok := <-c; !ok {
			return
		}
	}
}

// DeRegister 注销服务
func (r *SvcDisClient) DeRegister() error {
	r.lock.Lock()
	defer r.lock.Unlock()

	// lease revoke
	if _, err := r.cli.Revoke(r.ctx, r.leaseID); err != nil {
		return errors.Wrap(err, "租约解绑失败")
	}
	r.leaseID = 0
	return nil
}

func (r *SvcDisClient) Close() error {
	// close etcd client
	if err := r.cli.Close(); err != nil {
		return errors.Wrap(err, "关闭etcd连接失败")
	}
	return nil
}

func (r *SvcDisClient) HasLease() bool {
	return r.leaseID != 0
}
