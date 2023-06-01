package rpcpkg

import (
	"errors"
	"fmt"
	"net/rpc"
	"reflect"
)

type rpcTcpClient struct {
	cli *rpc.Client
}

func NewRpcTcpClient(address string) (*rpcTcpClient, error) {
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &rpcTcpClient{
		cli: client,
	}, nil
}

func (c *rpcTcpClient) GetReply(service any, method any, args any, reply any) error {
	rType := reflect.TypeOf(service).Elem()
	rVal := reflect.ValueOf(method)
	//rT := reflect.TypeOf(method)
	if reflect.TypeOf(method).Kind() != reflect.Func {
		return errors.New("rpc service method 需要传入 func")
	}
	_ = rVal

	fmt.Println(rType.Name())
	fmt.Println(fmt.Sprintf("%p", method))

	//serviceMethod := rType.Name()+"."+

	//for i := 0; i < rType.NumMethod(); i++ {
	//	fmt.Println(rType.Method(0))
	//}

	//addMethod, ok := rType.MethodByName("Add")
	//fmt.Println(addMethod, ok)
	return nil

	//var resp RpcResponse
	//if err := c.cli.Call(serviceMethod, args, &resp); err != nil {
	//	return errors.Wrap(err, "client call")
	//}
	//if resp.Err != nil {
	//	return resp.Err
	//}
	//
	//return json.Unmarshal(resp.Data, reply)
}
