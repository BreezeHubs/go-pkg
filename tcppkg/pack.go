package tcppkg

import (
	"encoding/binary"
	"io"
	"net"
)

var defaultHeaderLen int = 4

// Message structure for messages
type Message struct {
	DataLen uint32 // Length of the message
	Data    []byte // Content of the message
}

// Pack 封包方法,压缩数据
func Pack(msg []byte) []byte {
	buffer := make([]byte, defaultHeaderLen+len(msg))

	// 将 buffer 前面四个字节设置为包长度，大端序
	binary.BigEndian.PutUint32(buffer[0:4], uint32(len(msg)))
	copy(buffer[defaultHeaderLen:], msg)
	return buffer
}

// Unpack 拆包方法,解压数据
func Unpack(conn net.Conn) ([]byte, error) {
	var (
		msgLen uint32
		header = make([]byte, defaultHeaderLen)
	)
	if _, err := io.ReadFull(conn, header); err != nil {
		return nil, err
	}

	msgLen = binary.BigEndian.Uint32(header) // 转换成 10 进制的数字

	data := make([]byte, msgLen)
	if msgLen > 0 {
		_, e := io.ReadFull(conn, data) // 读取内容
		if e != nil {
			return nil, e
		}
	}

	return data, nil
}
