package client

import (
	"fmt"
	"testing"

	"github.com/xusworld/drpc/testdata"
)

func TestClient(t *testing.T) {
	optionFuncSet := []OptionFunc{
		WithService("Arith"),
		WithMethod("Mul"),
		WithNetwork("tcp"),
		WithAddr("127.0.0.1:10086"),
		WithProtocol("drpc"),
		WithSerializationType("Protobuf"),
		WithTimeout(DefaultReqTimeout),
		WithSendBuffSize(DefaultSendBuffSize),
		WithRecvBuffSize(DefaultRecvBuffSize),
	}
	c := NewClient(optionFuncSet)

	reply := &testdata.Reply{}

	_ = c.Call(&testdata.Args{
		Lhs: 3,
		Rhs: 4,
	}, reply)

	fmt.Println("The answer is ", reply.Ret)
	c.Stop()
}
