package server

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	optionFuncSet := []OptionFunc{
		WithServerAddress("127.0.0.1:8080"),
		WithServerNetwork("tcp"),
		WithProtocol("flash"),
		WithReadTimeout(time.Millisecond),
		WithWriteTimeout(time.Millisecond),
		WithSerializationType("protobuf"),
		WithKeepAlivePeriod(time.Millisecond),
	}

	options := &Options{}
	for _, optionFunc := range optionFuncSet {
		optionFunc(options)
	}

	assert.Equal(t, "127.0.0.1:8080", options.Address)
	assert.Equal(t, "tcp", options.Network)
	assert.Equal(t, "flash", options.Protocol)
	assert.Equal(t, time.Millisecond, options.ReadTimeout)
	assert.Equal(t, time.Millisecond, options.WriteTimeout)
	assert.Equal(t, "protobuf", options.SerializationType)
	assert.Equal(t, time.Millisecond, options.KeepAlivePeriod)
}
