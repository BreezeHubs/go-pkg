package tcppkg

import (
	"context"
	"fmt"
	"log"
	"net"
)

type tcpClient struct {
	ip string

	port int

	cli net.Conn
	ctx context.Context
}

func NewTcpClient(ctx context.Context, ip string, port int) (*tcpClient, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return nil, err
	}
	log.Println("[IP:" + conn.RemoteAddr().String() + "] 连接成功")

	return &tcpClient{
		ip:   ip,
		port: port,
		cli:  conn,
		ctx:  ctx,
	}, nil
}

func (c *tcpClient) Close() error {
	return c.cli.Close()
}
