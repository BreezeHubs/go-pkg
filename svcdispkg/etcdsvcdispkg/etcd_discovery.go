package etcdsvcdispkg

import (
	"context"
	"fmt"
	"github.com/BreezeHubs/go-pkg/svcdispkg"
	"github.com/BreezeHubs/go-pkg/typexpkg"
	"github.com/pkg/errors"
	clientv3 "go.etcd.io/etcd/client/v3"
	"math/rand"
	"sync/atomic"
)

type SvcDisServer struct {
	cli *clientv3.Client

	ctx context.Context
}

type EtcdSvcDisConfig = clientv3.Config

func NewSvcDisServer(ctx context.Context, conf *EtcdSvcDisConfig) (svcdispkg.IDiscovery, error) {
	if conf == nil || len(conf.Endpoints) == 0 {
		return nil, errors.New("endpoints is nil")
	}

	// 连接 etcd
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
		return nil, errors.Wrap(err, "连接etcd")
	}

	return &SvcDisServer{
		cli: cli,
		ctx: ctx,
	}, nil
}

func (d *SvcDisServer) GetServiceAddr(serviceName string, f func(i int64) int64) (string, error) {
	// get
	resp, err := d.cli.Get(d.ctx, serviceName, clientv3.WithPrefix())
	if err != nil {
		return "", errors.Wrap(err, "获取实例地址失败")
	}

	if len(resp.Kvs) == 0 {
		return "", errors.New("未找到可用实例地址")
	}

	// 负载均衡算法
	randIndex := f(int64(len(resp.Kvs)))
	return typexpkg.BytesToString(resp.Kvs[randIndex].Value), nil
}

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

	a := roundRobinLBIndex % i
	fmt.Println("A", a)
	return a
}

func (d *SvcDisServer) WatchService(serviceName string) error {
	//TODO implement me
	panic("implement me")
}
