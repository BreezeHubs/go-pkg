package tcppkg

import (
	"context"
	"fmt"
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
			bytes, err := ReadBytes(client.cli)
			log.Println("read: ", bytes, err)
			time.Sleep(1 * time.Second)
		}
	}()

	var count int
	for {
		md := []byte("哈哈哈哈哈哈哈哈哈哈哈哈 " + strconv.Itoa(count))
		fmt.Println("uint32(len(md))", len(md))
		err = SendBytes(client.cli, md)
		if err != nil {
			t.Log(err)
		}
		count++
		time.Sleep(1 * time.Second)
	}
}
