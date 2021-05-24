package client

import "time"

const (
	DefaultSerializationType = "protobuf"
	DefaultReqTimeout        = time.Millisecond * 42
	DefaultSendBuffSize      = 1024
	DefaultRecvBuffSize      = 1024
)

type Options struct {
	// Server service
	ServiceName string

	// Network
	Network string

	// Server address to connect to
	Addr string

	// Transport protocol
	Protocol string

	// Serialization Type
	// default serialization type is protobuf
	SerializationType string

	// Request timeout
	// default request timeout is 42ms
	Timeout time.Duration

	// Is request should be compressed
	// default NeedCompress is false which means not compress at all
	NeedCompress bool

	// Size of send buffer per each underlying connection in bytes
	SendBuffSize int

	// Size of recv buffer per each underlying connection in bytes
	RecvBuffSize int
}

type OptionFunc func(options *Options)

// ServiceName server service name
func WithServiceName(name string) OptionFunc {
	return func(options *Options) {
		options.ServiceName = name
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
