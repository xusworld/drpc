package client

import (
	"time"
)

const (
	DefaultSerializationType = "protobuf"
	DefaultReqTimeout        = time.Millisecond * 42
	DefaultSendBuffSize      = 1024
	DefaultRecvBuffSize      = 1024
)

// Options rpc client options
type Options struct {
	// Service name
	Service string

	// Method name
	Method string

	// Network type
	// Known networks are "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6",
	// "ip", "ip4", "ip6" and "unix"
	Network string

	// Server address to connect to
	// For TCP and UDP networks, the address has the form "host:port". The port must be a
	// literal port number or a service name.
	// Dial("tcp", "golang.org:http")
	// Dial("tcp", "192.0.2.1:http")
	// Dial("tcp", "198.51.100.1:80")
	// Dial("udp", "[2001:db8::1]:domain")
	// Dial("udp", "[fe80::1%lo0]:53")
	// Dial("tcp", ":80")
	//
	// For IP networks, the network must be "ip", "ip4" or "ip6" followed by a
	// colon and a literal protocol number or a protocol name, and the address has the form "host".
	// Dial("ip4:1", "192.0.2.1")
	// Dial("ip6:ipv6-icmp", "2001:db8::1")
	// Dial("ip6:58", "fe80::1%lo0")
	Addr string

	// Transport protocol
	// default protocol is "drpc"
	Protocol string

	// Serialization Type
	// Support Json, Protobuf, MessagePack and Thrift
	// default protocol is protobuf
	SerializationType string

	// Maximum request time
	// default request timeout is DefaultReqTimeout
	Timeout time.Duration

	concurrency int

	// Request compress flag
	// default NeedCompress is false which means not compress at all
	NeedCompress bool

	// Size of send buffer per each underlying connection in bytes
	SendBuffSize int

	// Size of recv buffer per each underlying connection in bytes
	RecvBuffSize int
}

type OptionFunc func(options *Options)

// WithService server service name
func WithService(name string) OptionFunc {
	return func(options *Options) {
		options.Service = name
	}
}

// WithMethod server service name
func WithMethod(name string) OptionFunc {
	return func(options *Options) {
		options.Method = name
	}
}

// WithNetwork network type
func WithNetwork(network string) OptionFunc {
	return func(options *Options) {
		options.Network = network
	}
}

// WithAddr server address to listen
func WithAddr(addr string) OptionFunc {
	return func(options *Options) {
		options.Addr = addr
	}
}

// WithProtocol rpc transport protocol
func WithProtocol(protocol string) OptionFunc {
	return func(options *Options) {
		options.Protocol = protocol
	}
}

// WithSerializationType serialization type
func WithSerializationType(typ string) OptionFunc {
	return func(options *Options) {
		options.SerializationType = typ
	}
}

// WithTimeout
func WithTimeout(timeout time.Duration) OptionFunc {
	return func(options *Options) {
		options.Timeout = timeout
	}
}

// WithNeedCompress
func WithNeedCompress(needCompress bool) OptionFunc {
	return func(options *Options) {
		options.NeedCompress = needCompress
	}
}

// WithSendBuffSize
func WithSendBuffSize(buffSize int) OptionFunc {
	return func(options *Options) {
		options.SendBuffSize = buffSize
	}
}

// WithRecvBuffSize
func WithRecvBuffSize(buffSize int) OptionFunc {
	return func(options *Options) {
		options.RecvBuffSize = buffSize
	}
}
