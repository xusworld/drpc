package client

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/xusworld/flash/codec"
	"github.com/xusworld/flash/protocol"
)

type clientCodec struct {
	options *Options
}

func NewClientCodec(options *Options) *clientCodec {
	return &clientCodec{
		options: options,
	}
}

func (cc *clientCodec) Encode(req interface{}) ([]byte, error) {
	// header
	header := &protocol.Header{}
	header.Magic = protocol.MagicNumber
	header.RequestId = 2020
	header.Priority = 42
	header.ServicePath = cc.options.Service
	header.ServiceMethod = cc.options.Method
	header.Timeout = int64(cc.options.Timeout)
	header.MessageType = protocol.MessageType_REQUEST
	header.CompressType = protocol.CompressType_GZIP
	header.SerializeType = serializeType2Int(cc.options.SerializationType)
	header.Status = make(map[string]string, 0)
	header.Status["user"] = "lukedong"

	// write body & body size
	body, err := codec.GetMarshaler(cc.options.SerializationType).Marshal(req)
	//body, err := marshaler.Marshal(req)
	if body == nil {
		fmt.Printf("body is nil , err is %s", err)
		return nil, errors.New("")
	}

	request := &protocol.Request{
		Header:  header,
		Payload: body,
	}

	// serialize request message
	buff, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}

	return buff, nil
}

func (cc *clientCodec) Decode(buff []byte) (*protocol.Response, error) {
	response := &protocol.Response{}

	err := proto.Unmarshal(buff, response)
	if err != nil {
		fmt.Printf("Decode err %s", err)
		return nil, err
	}
	return response, nil
}

func serializeType2Int(typ string) protocol.SerializeType {
	switch typ {
	case "Json":
		return protocol.SerializeType_JSON
	case "Protobuf":
		return protocol.SerializeType_PB
	case "Thrift":
		return protocol.SerializeType_THRIFT
	case "MessagePack":
		return protocol.SerializeType_MSGPACK
	default:
		return protocol.SerializeType_JSON
	}
}
