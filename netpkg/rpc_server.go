package netpkg

import (
	"bufio"
	"github.com/ugorji/go/codec"
	"io"
	"net"
	"net/rpc"
	"reflect"
	"sync"
)

func NewRpcServer(rpcAddr string, services ...any) error {
	// 新建rpc server
	server := rpc.NewServer()

	// 注册rpc对象
	//server.Register(new(service.RpcServer))
	for _, service := range services {
		if err := server.RegisterName(reflect.TypeOf(service).Elem().Name(), service); err != nil {
			return err
		}
	}

	laddr, err := net.ResolveTCPAddr("tcp", rpcAddr)
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		return err
	}

	// MessagePack: 是一种高效的二进制序列化格式。它允许你在多种语言(如JSON)之间交换数据。但它更快更小
	var mh codec.MsgpackHandle
	mh.MapType = reflect.TypeOf(map[string]interface{}(nil))

	var wg sync.WaitGroup
	for {
		wg.Add(1)
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}

		go func(conn *net.TCPConn) {
			defer wg.Done()

			// 用 buffer io做io解析提速
			bufConn := struct {
				io.Closer
				*bufio.Reader
				*bufio.Writer
			}{
				conn,
				bufio.NewReader(conn),
				bufio.NewWriter(conn),
			}
			server.ServeCodec(codec.MsgpackSpecRpc.ServerCodec(bufConn, &mh))
		}(conn)
	}
	wg.Wait()
	return nil
}
