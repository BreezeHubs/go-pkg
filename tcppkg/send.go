package tcppkg

import (
	"net"
)

func SendBytes(conn net.Conn, msgID uint16, msg []byte) error {
	bytes := Pack(msgID, msg)
	_, err := conn.Write(bytes)
	return err
}

//func SendProtobuf(conn net.Conn, msgID int, msg []byte) error {
//	bytes := Pack(msgID, msg)
//	_, err := conn.Write(bytes)
//	return err
//}
