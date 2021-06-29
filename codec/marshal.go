package codec

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/vmihailenco/msgpack"
)

const (
	Json         = "Json"
	Protobuf     = "Protobuf"
	Thrift       = "Thrift"
	Msgpack      = "MessagePack"
	Gob          = "Gob"
)

var (
	marshalers = make(map[string]Marshaler, 0)
)

func init() {
	marshalers[Json] = &JsonMarshaler{}
	marshalers[Protobuf] = &ProtobufMarshaler{}
	marshalers[Thrift] = &ThriftMarshaler{}
	marshalers[Msgpack] = &MsgpackMarshaler{}
}

// Marshaler encode/decode go struct
type Marshaler interface {
	// Marshal
	Marshal(v interface{}) ([]byte, error)

	// Unmarshal
	Unmarshal(data []byte, v interface{}) error
}

// RegisterMarshaller
func RegisterMarshaler(name string, marshaller Marshaler) {
	marshalers[name] = marshaller
}

func GetMarshaler(name string) Marshaler {
	return marshalers[name]
}

// JsonMarshaler
type JsonMarshaler struct{}

// Marshal
func (jm *JsonMarshaler) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal
func (jm *JsonMarshaler) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// ProtobufMarshaler
type ProtobufMarshaler struct{}

// Marshal
func (s *ProtobufMarshaler) Marshal(body interface{}) ([]byte, error) {
	msg, ok := body.(proto.Message)
	if !ok {
		return nil, errors.New("marshal fail: body not protobuf message")
	}
	return proto.Marshal(msg)
}

/// Unmarshal 反序列protobuf
func (s *ProtobufMarshaler) Unmarshal(in []byte, body interface{}) error {
	msg, ok := body.(proto.Message)
	if !ok {
		return errors.New("unmarshal fail: body not protobuf message")
	}
	return proto.Unmarshal(in, msg)
}


// ThriftMarshaler
type ThriftMarshaler struct{}

func (tm *ThriftMarshaler) Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	err := enc.Encode(v)
	return buf.Bytes(), err
}

func (tm *ThriftMarshaler) Unmarshal(data []byte, v interface{}) error {
	dec := msgpack.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(v)
	return err
}

// MsgpackMarshaler
type MsgpackMarshaler struct{}

func (mm *MsgpackMarshaler) Marshal(v interface{}) ([]byte, error) {
	return nil, nil
}

func (mm *MsgpackMarshaler) Unmarshal(data []byte, v interface{}) error {
	return nil
}
