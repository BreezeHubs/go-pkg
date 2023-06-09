package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/BreezeHubs/go-pkg/tcppkg"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	ml := make([]byte, 2)
	binary.LittleEndian.PutUint16(ml[0:2], uint16(1))
	fmt.Println("-----------uint16(len(msg)", ml)

	os.Exit(0)

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

	tcpServer := tcppkg.NewTcpServer(context.Background(), "127.0.0.1", 8911)
	log.Println("tcp server 启动成功")

	tcpServer.SetHookHandle(func(ctx context.Context, conn net.Conn) {
		for {
			if err := tcppkg.SendBytes(conn, 0, []byte("心跳")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	})

	//tcpServer.SetAcceptHandle(func(ctx context.Context, conn net.Conn) error {
	//	bytes, err := tcppkg.ReadBytes(conn)
	//	if err == nil {
	//		log.Println("接收数据：", string(bytes))
	//	}
	//	return err
	//})

	go func() {
		err := tcpServer.Start()
		log.Println(err)
	}()

	req()
}

func req() {
	fmt.Println("[]byte{0xaa, 0x55}", string([]byte{0xaa, 0x55}))

	ip := "127.0.0.1"
	port := 8911
	addr := "127.0.0.1:8911"
	client, err := tcppkg.NewTcpClient(context.Background(), ip, port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msgID, bytes, err := tcppkg.ReadBytes(client.GetConn())
		log.Println("-------client", msgID, string(bytes), err)
		if err == nil {
			log.Println(addr, "心跳检测：", msgID, string(bytes))
		}
		time.Sleep(1 * time.Second)
	}
}
