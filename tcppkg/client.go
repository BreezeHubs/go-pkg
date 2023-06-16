package tcppkg

import (
	"context"
	"fmt"
	"net"
	"time"
)

type tcpClient struct {
	ip string

	port int

	cli net.Conn
	ctx context.Context
}

func NewTcpClient(ctx context.Context, ip string, port int, t ...time.Duration) (*tcpClient, error) {
	timeout := 3 * time.Second
	if len(t) > 0 {
		timeout = t[0]
	}

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), timeout)
	if err != nil {
		return nil, err
	}

	return &tcpClient{
		ip:   ip,
		port: port,
		cli:  conn,
		ctx:  ctx,
	}, nil
}

func (c *tcpClient) GetConn() net.Conn {
	return c.cli
}

func (c *tcpClient) Close() error {
	return c.cli.Close()
}
