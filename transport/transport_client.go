package transport

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/xusworld/flash/config"
	"github.com/xusworld/flash/errors"
	"github.com/xusworld/flash/log"
)

// ClientTransport transport bytes stream from client to server
func ClientTransport(ctx context.Context, buf []byte, opts *Options) ([]byte, error) {
	defaultTrans := clientTransporter{}

	rsp, err := defaultTrans.Transport(ctx, buf, opts)

	if err != nil {
		log.Debugf("Transport err %s", err)
	}
	return rsp, err
}

// clientTransporter client transport implement
type clientTransporter struct {
	Transporter
}

// Transport bytes stream from client to server and get reply from server
func (ct *clientTransporter) Transport(ctx context.Context, buf []byte, opts *Options) ([]byte, error) {
	switch opts.Network {
	case "tcp":
		return ct.tcpTransport(ctx, buf, opts)
	case "udp":
		return ct.udpTransport(ctx, buf, opts)
	default:
		return nil, errors.UnknownNetworkType
	}
}

// tcpTransport
func (ct *clientTransporter) tcpTransport(ctx context.Context, buff []byte, opts *Options) ([]byte, error) {
	var conn net.Conn
	var err error
	dialChan := make(chan struct{})

	// Start a goroutine to connects the server
	go func() {
		conn, err = net.Dial(opts.Network, opts.Addr)
		close(dialChan)
	}()

	select {
	case <-ctx.Done():
		log.Debug(config.Project + ".Client Context Timeout")
		return nil, errors.ClientContextTimeout
	case <-dialChan:
		if err != nil {
			log.Debugf(config.Project+". call net.Dial() error %s", err)
			return nil, err
		}
	}

	reply := tcpConnectHandler(ctx, conn, buff)

	return reply, nil
}

// tcpConnectHandler
func tcpConnectHandler(ctx context.Context, conn net.Conn, buff []byte) (resp []byte) {
	defer conn.Close()

	var writeErr error
	var readErr error
	readChan := make(chan struct{})

	go func() {
		writeErr = Write(conn, buff)

		if writeErr != nil {
			log.Debug("flash: tcpTransport wirte error")
		}
	}()

	go func() {
		resp, readErr = ioutil.ReadAll(conn)
		if readErr != nil {
			log.Debug("err %s", readErr)
		}
		close(readChan)
	}()

	select {
	case <-ctx.Done():
		log.Debug("context timeout")
		return nil
	case <-readChan:
		if readErr != nil {
			log.Debugf("tcpConnectHandler read error %s", readErr)
			return nil
		}
		return resp
	}
}

// udpTransport
func (ct *clientTransporter) udpTransport(ctx context.Context, buf []byte, opts *Options) ([]byte, error) {
	return nil, nil
}

func Write(conn net.Conn, buff []byte) error {
	totalBytes := len(buff)
	sentBytes := 0

	for sentBytes < totalBytes {
		n, err := conn.Write(buff)

		if err != nil {
			sentBytes += n
			fmt.Printf("write %d bytes, error: %s\n", n, err)
			return err
		}

		sentBytes += n
	}

	return nil
}
