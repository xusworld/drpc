package transport

import (
	"context"
	"net"
)

// Transporter transport interface for server and client
type Transporter interface {
	Read(conn net.Conn) ([]byte, error)

	Write(context.Context, []byte, Options) ([]byte, error)
}
