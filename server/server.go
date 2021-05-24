package server

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/xusworld/crpc/log"
	"github.com/xusworld/crpc/transport"
	"github.com/xusworld/crpc/util"
)

type Server struct {
	options  *Options
	listener transport.Listener

	sync.RWMutex
}

func NewServer(optionFuncSet []OptionFunc) *Server {
	server := Server{}

	// server options
	server.options = &Options{}
	for _, optionFunc := range optionFuncSet {
		optionFunc(server.options)
	}

	// network type
	// TODO util.Contains 性能比较差
	if util.Contains([]string{"tcp", "tcp4", "tcp6"}, server.options.Network) {
		server.listener = &transport.TcpListener{}
		_ = server.listener.Init(server.options.Address)
	} else {
		// UDP server
	}

	return &server
}

func (s *Server) Serve() {
	defer s.listener.Close()

	s.listener.Accept(context.Background(), serve)
}

func serve(ctx context.Context, conn net.Conn) {
	buff := make([]byte, 1024)
	num, err := conn.Read(buff)
	if err != nil {
		log.Error(err.Error())
	}
	msg := string(buff[:num])
	fmt.Println("recv msg", msg)

	conn.Write([]byte("hello world from server"))
}

func (s *Server) Addr() string {
	return s.listener.Addr()
}
