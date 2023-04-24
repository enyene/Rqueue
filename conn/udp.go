package conn

import (
	"net"
)

// writer for sending udp mesaages
type UdpWriter struct {
	address    *net.UDPAddr
	connection *net.UDPConn
}

// newudep writer

func NewUDPWriter(connect *net.UDPConn, addr *net.UDPAddr) *UdpWriter {
	return &UdpWriter{
		address:    addr,
		connection: connect,
	}
}

func (uw *UdpWriter) Write(b []byte) (int, error) {
	return uw.connection.WriteToUDP(b, uw.address)
}
