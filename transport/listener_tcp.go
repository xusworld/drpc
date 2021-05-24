package transport

import (
	"context"
	"net"

	"github.com/xusworld/crpc/log"
)

// TcpListener
type TcpListener struct {
	listener net.Listener
}

type TcpHandler func(ctx context.Context, conn net.Conn)

func (tl *TcpListener) Init(addr string) error {
	listener, err := net.Listen(tcp, addr)

	if err != nil {
		log.Error(err.Error())
	}

	l, ok := listener.(*net.TCPListener)

	if !ok {

	}
	tl.listener = l

	return nil
}

func (tl *TcpListener) Accept(ctx context.Context, handler TcpHandler) (net.Conn, error) {

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:

		}

		conn, err := tl.listener.Accept()
		if err != nil {
			log.Error()
		}

		go handler(ctx, conn)
	}
}

func (tl *TcpListener) Close() error {
	err := tl.listener.Close()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (tl *TcpListener) Addr() string {
	return tl.Addr()
}

func tcpHandler(ctx context.Context, conn net.Conn) {
	defer conn.Close()

	for {
		select {
		case <-ctx.Done():
			return
		default:



		}
	}
}


