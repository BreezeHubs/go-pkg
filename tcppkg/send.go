package tcppkg

import (
	"net"
)

func SendBytes(conn net.Conn, msg []byte) error {
	bytes := Pack(msg)
	_, err := conn.Write(bytes)
	return err
}

func SendProtobuf(conn net.Conn, msg []byte) error {
	bytes := Pack(msg)
	_, err := conn.Write(bytes)
	return err
}
