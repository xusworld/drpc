package client

import (
	"sync"
)

// RPCClient RPC client interface
type RPCClient interface {
	Start()

	Stop()

	Call(request interface{}) (interface{}, error)

	Send(request interface{}) error
}

// Client RPC client
type Client struct {
	options *Options
	mutex   sync.RWMutex
}

func NewClient(optionFuncSet []OptionFunc) *Client {

	client := &Client{}

	for _, optionFunc := range optionFuncSet {
		optionFunc(client.options)
	}

	setDefaultIfNecessary(client.options)

	return client
}

// Start
func (c *Client) Start() {

}

// Stop
func (c *Client) Stop() {

}

// call
func (c *Client) Call(request interface{}) (interface{}, error) {

	return struct{}{}, nil
}

// Send
func (c *Client) Send(request interface{}) error {
	return nil
}

// setDefaultIfNecessary
func setDefaultIfNecessary(options *Options) {
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
