package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	optionFuncSet := []OptionFunc{
		WithService("Proxy"),
		WithMethod("Call"),
		WithNetwork("tcp"),
		WithAddr("127.0.0.1:8080"),
		WithProtocol("drpc"),
		WithSerializationType(DefaultSerializationType),
		WithTimeout(DefaultReqTimeout),
		WithSendBuffSize(DefaultSendBuffSize),
		WithRecvBuffSize(DefaultRecvBuffSize),
	}

	options := &Options{}

	for _, optionFunc := range optionFuncSet {
		optionFunc(options)
	}

	assert.Equal(t, "Proxy", options.Service)
	assert.Equal(t, "Call", options.Method)
	assert.Equal(t, "tcp", options.Network)
	assert.Equal(t, "127.0.0.1:8080", options.Addr)
	assert.Equal(t, "drpc", options.Protocol)
	assert.Equal(t, DefaultSerializationType, options.SerializationType)
	assert.Equal(t, DefaultReqTimeout, options.Timeout)
	assert.Equal(t, DefaultSendBuffSize, options.SendBuffSize)
	assert.Equal(t, DefaultRecvBuffSize, options.RecvBuffSize)
}
