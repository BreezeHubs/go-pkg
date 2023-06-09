package timepkg

// NewRetry 重试任务
// @param f func()error 重试的任务，返回error代表重试
// @param times int 重试次数
// @param t time.Duration 重试时间间隔
// @param timeout time.Duration 每次任务超时时间，会计入重试
//func NewRetry(f func() error, times int, t time.Duration, timeout time.Duration) error {
//	var (
//		errC = make(chan error, 1)
//		e    error
//	)
//	for {
//		if times <= 0 {
//			return e
//		}
//
//		go func() {
//			errC <- f()
//		}()
//		select {
//		case err := <-errC:
//			if err != nil {
//				e = err
//				times--
//			}
//		case <-time.After(timeout):
//			e = errors.New("timeout")
//		}
//	}
//}
