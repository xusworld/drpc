package transport

import (
	"context"
	"net"
	"time"
)

// Options transport options for server and client
type Options struct {
	Network        string
	Addr           string
	Protocol       string
	Timeout        time.Duration
	tcpConnHandler TcpConnHandler
	udpConnHandler UdpConnHandler
}

// OptionFunc
type OptionFunc func(options *Options)

// TcpConnHandler tcp connection handler
type TcpConnHandler func(ctx context.Context, conn net.Conn, options *Options) error

// UdpConnHandler udp connection handler
type UdpConnHandler func(ctx context.Context, conn net.Conn, options *Options) error

// WithNetwork function closures for set transport options "network"
func WithNetwork(network string) OptionFunc {
	return func(options *Options) {
		options.Network = network
	}
}

// WithAddr function closures for set transport options "addr"
func WithAddr(addr string) OptionFunc {
	return func(options *Options) {
		options.Addr = addr
	}
}
