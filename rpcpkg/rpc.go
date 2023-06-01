package rpcpkg

import (
	"net"
	"net/rpc"
)

type rpcServer struct {
	svr []any
}

func NewRpcServer() *rpcServer {
	return &rpcServer{}
}

func (s *rpcServer) Register(svr ...any) *rpcServer {
	s.svr = append(s.svr, svr...)
	return s
}

func (s *rpcServer) TcpListen(address string, acceptErrorFunc func(err error)) error {
	for _, item := range s.svr {
		if err := rpc.Register(item); err != nil {
			return err
		}
	}

	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			if acceptErrorFunc != nil {
				acceptErrorFunc(err)
			}
		}
		go rpc.ServeConn(conn)
	}
}
