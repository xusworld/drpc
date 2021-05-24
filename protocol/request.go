package protocol

import (
	"github.com/xusworld/crpc/client"
	"github.com/xusworld/crpc/codec"
)

const (
	DefaultCrcCode = 42
)
// Request rpc request data structure
type Request struct {
	// request header
	header RequestHeader

	// request body
	body RequestBody
}

// RequestHead
type RequestHeader struct {
	// CRC code
	CrcCode int

	// Message length, include header and body
	Length int

	// 0 业务消息，1 业务响应消息，2 业务 ONE WAY 消息，3 握手请求消息
	// 4 握手应答消息，5 心跳请求消息，6 心跳应答消息
	Type uint8

	// 消息优先级
	Priority uint8

	// 请求服务路径
	ServicePath string

	// 请求服务方法
	ServiceMethod string

	// 透传数据
	Attachment map[string]string
}

type RequestBody []byte

func BuildRequest(request interface{}, options *client.Options) *Request {
	header := RequestHeader{}
	header.CrcCode = DefaultCrcCode
	header.Length = 1
	header.Type = 1
	header.Priority = 1
	header.ServicePath = options.ServiceName
	header.ServiceMethod = options.ServiceName
	header.Attachment = make(map[string]string, 0)

	codec := codec.Codec()

}
