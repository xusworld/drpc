package server

import "net"

const (
	ReadBuffSize = 1024
)

// Server interface
// 1. tcp tcpServer
// 2. udp tcpServer
// 3. http tcpServer
// 其他特性server，如性能优化
type Server interface {
	Serve() error

	Stop() error

	Address() net.Addr
}
