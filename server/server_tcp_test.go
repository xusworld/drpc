package server

import (
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	optionFuncSet := []OptionFunc{
		WithServerAddress("127.0.0.1:10086"),
		WithServerNetwork("tcp"),
		WithProtocol("drpc"),
		WithReadTimeout(time.Millisecond),
		WithWriteTimeout(time.Millisecond),
		WithSerializationType("Protobuf"),
		WithKeepAlivePeriod(time.Millisecond),
	}
	server := NewTcpServer(optionFuncSet)
	service := &defaultService{}
	_ = service.Register(new(Arith))
	_ = server.Register(service)

	_ = server.Serve()
}
