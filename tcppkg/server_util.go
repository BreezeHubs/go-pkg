package tcppkg

import (
	"log"
	"net"
	"time"
)

func (s *tcpServer) ipListConn(conn net.Conn, err error) {
	if err == nil {
		s.ipListLock.Lock()
		defer s.ipListLock.Unlock()

		_, ok := s.ipList[conn.RemoteAddr().String()]
		if !ok {
			s.ipList[conn.RemoteAddr().String()] = struct{}{}
			log.Println("[IP:" + conn.RemoteAddr().String() + "] 已连接")

			if s.hookHandle != nil {
				if err = s.hookHandle(s.ctx, conn); err != nil {
					return
				} // 处理首次连接钩子
			}
		}
	}
}

func (s *tcpServer) connHandle(conn net.Conn) {
	defer conn.Close()

	for {
		if s.acceptHandle != nil {
			if err := s.acceptHandle(s.ctx, conn); err != nil {
				return
			}
		}
		time.Sleep(200 * time.Microsecond)
	}
}
