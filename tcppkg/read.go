package tcppkg

import (
	"net"
)

func ReadBytes(conn net.Conn) (uint16, []byte, error) {
	return Unpack(conn)
}

//func ReadProtobuf(conn net.Conn) (int, []byte, error) {
//	return Unpack(conn)
//}
