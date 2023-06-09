package tcppkg

//type RouterFunc func(data []byte) error
//
//func (s *tcpServer) AddRouter(msgID uint16, f RouterFunc) {
//	s.routeLock.Lock()
//	defer s.routeLock.Unlock()
//
//	s.route[msgID] = f
//}

//func (s *tcpServer) EffectRoute(df ...func(conn net.Conn)) {
//	var dFunc func(conn net.Conn)
//	if len(df) > 0 {
//		dFunc = df[0]
//	}
//
//	s.SetAcceptHandle(func(ctx context.Context, conn net.Conn) bool {
//		msgID, bytes, err := ReadBytes(conn)
//		if err != nil {
//			return err
//		}
//
//		s.routeLock.RLock()
//		f, ok := s.route[msgID]
//		s.routeLock.RUnlock()
//		if !ok {
//			return err
//		}
//
//		return f(bytes)
//	}, dFunc)
//}
