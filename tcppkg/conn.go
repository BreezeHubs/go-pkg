package tcppkg

import (
	"net"
)

func (s *tcpServer) handleConn(conn net.Conn) {
	s.ipListConn(conn) // 维护IP列表、处理首次连接钩子
	s.connHandle(conn)
}

// 保存 ip 到 list，并执行首次连接的钩子
func (s *tcpServer) ipListConn(conn net.Conn) {
	s.ipListLock.Lock()
	defer s.ipListLock.Unlock()

	if _, ok := s.ipList[conn.RemoteAddr().String()]; !ok {
		s.ipList[conn.RemoteAddr().String()] = struct{}{}

		if s.hookHandle != nil {
			s.hookHandle(s.ctx, conn) // 处理首次连接钩子
		}
	}
}

type ConnHandleResult int

const (
	ConnContinue ConnHandleResult = iota
	ConnStop
)

// 监听 conn
func (s *tcpServer) connHandle(conn net.Conn) {
	defer func() {
		conn.Close()
		s.ipListLock.Lock()
		delete(s.ipList, conn.RemoteAddr().String())
		s.ipListLock.Unlock()
		if s.acceptHandleDeferFunc != nil {
			s.acceptHandleDeferFunc(conn)
		}
	}()

	for {
		if s.acceptHandle != nil {
			result := s.acceptHandle(s.ctx, conn)
			switch result {
			case ConnContinue:
				continue
			case ConnStop:
				return
			default:
			}
		}
	}
}
