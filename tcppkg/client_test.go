package tcppkg

import (
	"context"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestNewTcpClient(t *testing.T) {
	client, err := NewTcpClient(context.Background(), "127.0.0.1", 8900)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		for {
			_, bytes, err := ReadBytes(client.cli)
			if err == nil {
				log.Println("接收数据：", string(bytes))
			}
			time.Sleep(1 * time.Second)
		}
	}()

	var count int
	for {
		md := []byte("哈哈哈哈哈哈哈哈哈哈哈哈 " + strconv.Itoa(count))
		err = SendBytes(client.cli, 0, md)
		if err != nil {
			t.Log(err)
		}
		count++
		time.Sleep(1 * time.Second)
	}
}
