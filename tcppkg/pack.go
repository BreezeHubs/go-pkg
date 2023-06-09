package tcppkg

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	"reflect"
)

var (
	headLen          = 2
	msgIDLen         = 2
	defaultHeaderLen = 2
)

// Message structure for messages
type Message struct {
	HeadLen uint32
	DataLen uint32 // Length of the message
	Data    []byte // Content of the message
}

// Pack 封包方法,压缩数据
func Pack(msgID uint16, msg []byte) []byte {
	buffer := make([]byte, headLen+msgIDLen+defaultHeaderLen+len(msg))

	// 将 buffer 前面2个字节设置为包头
	buffer[0], buffer[1] = 0xaa, 0x55

	// 将 buffer 2个字节设置为包长度，大端序 msg id
	binary.LittleEndian.PutUint16(buffer[headLen:headLen+msgIDLen], msgID)

	// 将 buffer 前面四个字节设置为包长度，大端序
	binary.LittleEndian.PutUint16(buffer[headLen+msgIDLen:headLen+msgIDLen+defaultHeaderLen], uint16(len(msg)))

	copy(buffer[headLen+msgIDLen+defaultHeaderLen:], msg)
	return buffer
}

// Unpack 拆包方法,解压数据
func Unpack(conn net.Conn) (uint16, []byte, error) {
	var (
		headHeader    = make([]byte, headLen)
		msgHeader     = make([]byte, msgIDLen)
		dataLenHeader = make([]byte, defaultHeaderLen)
	)

	if _, err := conn.Read(headHeader); err != nil {
		return 0, nil, err
	}
	if !reflect.DeepEqual(headHeader, []byte{0xaa, 0x55}) {
		return 0, nil, errors.New("head error")
	}

	if _, err := io.ReadFull(conn, msgHeader); err != nil {
		return 0, nil, err
	}
	msgID := binary.LittleEndian.Uint16(msgHeader) // 转换成 10 进制的数字

	if _, err := io.ReadFull(conn, dataLenHeader); err != nil {
		return 0, nil, err
	}
	dataLen := binary.LittleEndian.Uint16(dataLenHeader) // 转换成 10 进制的数字

	data := make([]byte, dataLen)
	if dataLen > 0 {
		_, e := io.ReadFull(conn, data) // 读取内容
		if e != nil {
			return 0, nil, e
		}
	}

	// fmt.Println("------data", headHeader, msgHeader, dataLenHeader, data)
	return msgID, data, nil
}
