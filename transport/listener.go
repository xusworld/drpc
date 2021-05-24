package transport

import (
	"context"
	"net"
)

const (
	tcp = "tcp"
	udp = "udp"
)

// Listener is network listener for stream-oriented protocols
type Listener interface {
	Init(addr string) error

	Accept(ctx context.Context, handler TcpHandler) (net.Conn, error)

	Close() error

	Addr() string
}

