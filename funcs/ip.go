package funcs

import (
	"encoding/binary"
	"net"
)

// Ip2long ip转换成uint
func Ip2long(ipAddr string) uint32 {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip.To4())
}

// Long2ip uint to ip
func Long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}
