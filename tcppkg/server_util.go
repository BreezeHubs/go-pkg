package tcppkg

import (
	"net"
	"time"
)

func (s *tcpServer) ipListConn(conn net.Conn) {
	s.ipListLock.Lock()
	defer s.ipListLock.Unlock()

	_, ok := s.ipList[conn.RemoteAddr().String()]
	if !ok {
		s.ipList[conn.RemoteAddr().String()] = struct{}{}

		if s.hookHandle != nil {
			s.hookHandle(s.ctx, conn) // 处理首次连接钩子
		}
	}
}

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
			if !s.acceptHandle(s.ctx, conn) {
				return
			}
		}
		time.Sleep(200 * time.Microsecond)
	}
}
