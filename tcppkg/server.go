package tcppkg

import (
	"context"
	"fmt"
	"net"
	"sync"
)

type (
	TcpServerOption        func(s *tcpServer)
	HandleFuncWithContinue func(ctx context.Context, conn net.Conn) ConnHandleResult
	HandleFunc             func(ctx context.Context, conn net.Conn)

	tcpServer struct {
		ip string

		port int

		svr net.Listener

		acceptHandle          HandleFuncWithContinue
		acceptHandleDeferFunc func(conn net.Conn)
		hookHandle            HandleFunc

		ctx context.Context

		ipList     map[string]struct{}
		ipListLock sync.Mutex

		//route     map[uint16]RouterFunc
		//routeLock sync.RWMutex
	}
)

func NewTcpServer(ctx context.Context, ip string, port int, opt ...TcpServerOption) *tcpServer {
	return &tcpServer{
		ip:     ip,
		port:   port,
		ctx:    ctx,
		ipList: make(map[string]struct{}, 6),
		//route:  make(map[uint16]RouterFunc, 12),
	}
}

func (s *tcpServer) SetAcceptHandle(f HandleFuncWithContinue, df func(conn net.Conn)) {
	s.acceptHandle = f
	s.acceptHandleDeferFunc = df
}

func (s *tcpServer) SetHookHandle(f HandleFunc) {
	s.hookHandle = f
}

func (s *tcpServer) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		return err
	}
	//s.svr = listener
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go s.handleConn(conn)
	}
	return nil
}

func (s *tcpServer) conn() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		return err
	}
	s.svr = listener
	return nil
}

func (s *tcpServer) Close() error {
	return s.svr.Close()
}
