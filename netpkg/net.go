package netpkg

import (
	"net"
	"strings"
)

func GetLocalIp() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "", err
	}

	localIp := strings.Split(conn.LocalAddr().String(), ":")[0]
	conn.Close()

	return localIp, nil
}
