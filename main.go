package main

import (
	"context"
	"github.com/BreezeHubs/go-pkg/tcppkg"
	"log"
	"net"
)

func main() {
	//s := schedpkg.NewSchedule(context.Background())
	//s.Add("", func(ctx context.Context) error {
	//	for true {
	//
	//	}
	//	return nil
	//})
	//
	//err := s.RunAndGracefullyExit()
	//log.Println(err)

	tcpServer, err := tcppkg.NewTcpServer(context.Background(), "127.0.0.1", 8900)
	if err != nil {
		panic(err)
	}
	log.Println("tcp server 启动成功")

	//tcpServer.SetHookHandle(func(ctx context.Context, conn net.Conn) error {
	//	for {
	//		if err = tcpServer.Send(conn, []byte("心跳")); err != nil {
	//			return err
	//		}
	//		time.Sleep(1 * time.Second)
	//	}
	//})

	tcpServer.SetAcceptHandle(func(ctx context.Context, conn net.Conn) error {
		bytes, err := tcppkg.ReadBytes(conn)
		log.Println(string(bytes), err)
		return err
	})

	err = tcpServer.Start()
	log.Println(err)
}
