package server

import (
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	optionFuncSet := []OptionFunc{
		WithServerAddress("127.0.0.1:10086"),
		WithServerNetwork("tcp"),
		WithProtocol("crpc"),
		WithReadTimeout(time.Millisecond),
		WithWriteTimeout(time.Millisecond),
		WithSerializationType("protobuf"),
		WithKeepAlivePeriod(time.Millisecond),
	}
	server := NewServer(optionFuncSet)
	server.Serve()
}
