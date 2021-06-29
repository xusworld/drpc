package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/xusworld/flash/testdata"
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *testdata.Args, reply *testdata.Reply) error {
	reply.Ret = args.Lhs * args.Rhs
	return nil
}

func TestService(t *testing.T) {
	s := defaultService{}
	_ = s.Register(new(Arith))

	args := new(testdata.Args)
	args.Lhs = 21
	args.Rhs = 2

	reply := new(testdata.Reply)

	fmt.Println("status", s.status)
	_ = s.Call("Mul", context.Background(), args, reply)
	fmt.Println("reply ", reply.Ret)
}
