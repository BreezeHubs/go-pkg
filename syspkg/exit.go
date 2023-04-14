package syspkg

import (
	"os"
	"os/signal"
	"syscall"
)

type eSignal struct {
	signal chan os.Signal
}

func NewListenExitSignal() *eSignal {
	// 从这里开始优雅退出监听系统信号，强制退出以及超时强制退出。
	e := &eSignal{
		signal: make(chan os.Signal, 1),
	}

	//windows
	signal.Notify(e.signal, os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGHUP,
		syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP,
		syscall.SIGABRT, syscall.SIGTERM)

	//linux & mac
	//signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGKILL, syscall.SIGSTOP,
	//	syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL,
	//	syscall.SIGTRAP, syscall.SIGABRT, syscall.SIGSYS, syscall.SIGTERM)

	return e
}

func (e *eSignal) Wait() {
	select {
	case <-e.signal:
		go func() {
			select {
			case <-e.signal:
				os.Exit(1) //再次监听退出信号
			}
		}()
	}
}

func (e *eSignal) Signal() <-chan os.Signal {
	return e.signal
}
