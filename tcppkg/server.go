package tcppkg

import (
	"context"
	"fmt"
	"net"
	"sync"
)

type (
	TcpServerOption func(s *tcpServer)
	HandleFunc      func(ctx context.Context, conn net.Conn) error

	tcpServer struct {
		ip string

		port int

		svr net.Listener

		acceptHandle HandleFunc
		hookHandle   HandleFunc

		ctx context.Context

		ipList     map[string]struct{}
		ipListLock sync.Mutex
	}
)

func NewTcpServer(ctx context.Context, ip string, port int, opt ...TcpServerOption) (*tcpServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return nil, err
	}

	return &tcpServer{
		ip:     ip,
		port:   port,
		svr:    listener,
		ctx:    ctx,
		ipList: make(map[string]struct{}, 6),
	}, nil
}

func (s *tcpServer) SetAcceptHandle(f HandleFunc) {
	s.acceptHandle = f
}

func (s *tcpServer) SetHookHandle(f HandleFunc) {
	s.hookHandle = f
}

func (s *tcpServer) Start() error {
	for {
		conn, err := s.svr.Accept()

		s.ipListConn(conn, err) // 维护IP列表、处理首次连接钩子
		go s.connHandle(conn)
	}
}

func (s *tcpServer) Close() error {
	return s.svr.Close()
}
