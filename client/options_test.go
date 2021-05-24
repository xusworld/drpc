package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	optionFuncSet := []OptionFunc{
		WithServiceName("Add"),
		WithNetwork("tcp"),
		WithAddr("127.0.0.1:8080"),
		WithProtocol("crpc"),
		WithSerializationType("protobuf"),
		WithTimeout(DefaultReqTimeout),
		WithNeedCompress(DefaultNeedCompress),
		WithSendBuffSize(DefaultSendBuffSize),
		WithRecvBuffSize(DefaultRecvBuffSize),
	}

	options := &Options{}

	for _, optionFunc := range optionFuncSet {
		optionFunc(options)
	}

	assert.Equal(t, "Add", options.ServiceName)
	assert.Equal(t, "tcp", options.Network)
	assert.Equal(t, "127.0.0.1:8080", options.Addr)
	assert.Equal(t, "crpc", options.Protocol)
	assert.Equal(t, DefaultSerializationType, options.SerializationType)
	assert.Equal(t, DefaultReqTimeout, options.Timeout)
	assert.Equal(t, DefaultNeedCompress, options.NeedCompress)
	assert.Equal(t, DefaultSendBuffSize, options.SendBuffSize)
	assert.Equal(t, DefaultRecvBuffSize, options.RecvBuffSize)
}
