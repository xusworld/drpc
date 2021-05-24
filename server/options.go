package server

import (
	"time"
)

// Options server options
type Options struct {
	Address           string
	Network           string
	Protocol          string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SerializationType string
	KeepAlivePeriod   time.Duration
}

type OptionFunc func(options *Options)

// WithServerAddress address is a IP:Port format string
func WithServerAddress(address string) OptionFunc {
	return func(options *Options) {
		options.Address = address
	}
}

// WithServerNetwork network is none of tcp, tcp4, tcp6, udp, udp4 and udp6
func WithServerNetwork(network string) OptionFunc {
	return func(options *Options) {
		options.Network = network
	}
}

// WithProtocol protocol
func WithProtocol(protocol string) OptionFunc {
	return func(options *Options) {
		options.Protocol = protocol
	}
}

// WithReadTimeout RPC server read timeout
func WithReadTimeout(timeout time.Duration) OptionFunc {
	return func(options *Options) {
		options.ReadTimeout = timeout
	}
}

// WithWriteTimeout RPC server write timeout
func WithWriteTimeout(timeout time.Duration) OptionFunc {
	return func(options *Options) {
		options.WriteTimeout = timeout
	}
}

// WithSerialization serializationType is one of json, protobuf, messagepack
func WithSerializationType(serializationType string) OptionFunc {
	return func(options *Options) {
		options.SerializationType = serializationType
	}
}

// WithKeepAlivePeriod
func WithKeepAlivePeriod(keepAlivePeriod time.Duration) OptionFunc {
	return func(options *Options) {
		options.KeepAlivePeriod = keepAlivePeriod
	}
}
