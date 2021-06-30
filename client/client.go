package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/xusworld/drpc/codec"
	"github.com/xusworld/drpc/errors"
	"github.com/xusworld/drpc/log"
	"github.com/xusworld/drpc/transport"
)

// Client RPC client interface
type Client interface {
	// check
	Start() error

	// Async call  server
	Call(args interface{}, reply interface{}) error

	// send one-way message to server, no need to get any reply from server
	Send(args interface{}) error

	// stop client and close all channels and network connection
	Stop()
}

// Client RPC client
type client struct {
	// rpc client options
	options *Options
	// wait group
	waitGroup sync.WaitGroup
	//
	ctx        context.Context
	cancelFunc context.CancelFunc

	// done signal
	done chan struct{}
	// stop signal
	stop chan struct{}

	mutex sync.RWMutex
}

// NewClient create a rpc client
func NewClient(optionFuncSet []OptionFunc) Client {
	client := &client{
		options: &Options{},
	}

	// function closures to set this client's all options
	for _, optionFunc := range optionFuncSet {
		optionFunc(client.options)
	}

	// set this client's all default options
	setDefaultOptions(client.options)

	return client
}

func (c *client) Start() error {
	if c.options.Timeout <= 0 {
		log.Debugf("Client option timeout is %d, less than 0", c.options.Timeout)
		return errors.ClientOptionError
	}

	c.ctx, c.cancelFunc = context.WithTimeout(context.Background(), c.options.Timeout)

	select {
	case <-c.ctx.Done():
		return errors.ClientContextTimeout
	default:
	}

	return nil
}

// CallTimeout
func (c *client) Call(args, reply interface{}) (err error) {
	c.ctx = context.Background()
	return c.call(args, reply)
}

// call
func (c *client) call(args, reply interface{}) error {

	// client codec
	cc := NewClientCodec(c.options)
	request, _ := cc.Encode(args)

	// setting transport.Options
	options := &transport.Options{
		Network: c.options.Network,
		Addr:    c.options.Addr,
		Timeout: c.options.Timeout,
	}

	// call server and get reply from server
	buff, err := transport.ClientTransport(c.ctx, request, options)
	if err != nil {
		fmt.Printf("drpc: ClientTransport error %s", err)
	}

	// parse protocol header
	response, err := cc.Decode(buff)
	if err != nil {
		log.Debug("client.Decode error %s", err)
	}

	serializer := codec.GetMarshaler(c.options.SerializationType)
	err = serializer.Unmarshal(response.Payload, reply)

	return err
}

// Send
func (c *client) Send(args interface{}) error {
	return nil
}

// Stop
func (c *client) Stop() {
	c.waitGroup.Wait()
}

// setDefaultOptions set default and necessary options
func setDefaultOptions(options *Options) {
	if options.SerializationType == "" {
		options.SerializationType = DefaultSerializationType
	}

	if options.Timeout == 0 {
		options.Timeout = DefaultReqTimeout
	}

	if options.SendBuffSize == 0 {
		options.SendBuffSize = DefaultSendBuffSize
	}

	if options.RecvBuffSize == 0 {
		options.RecvBuffSize = DefaultRecvBuffSize
	}
}
