package tcppkg

import (
	"net"
)

func ReadBytes(conn net.Conn) ([]byte, error) {
	return Unpack(conn)
}

func ReadProtobuf(conn net.Conn) ([]byte, error) {
	return Unpack(conn)
}
