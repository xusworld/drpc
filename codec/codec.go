package codec

import (
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
	pb "google.golang.org/protobuf/proto"
)

// Codec Interface
type Codec interface {
	// Encode
	Encode(v interface{}) ([]byte, error)

	// Decode
	Decode(data []byte, v interface{}) error
}

// JsonCodec json codec implement
type JsonCodec struct {
}

// Encode return the JSON encoding of v
func (c *JsonCodec) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Decode parses the JSON-encoded data and stores the result in the value
// pointed to by v. If v is nil or not a pointer, returns an InvalidUnmarshalError
func (c *JsonCodec) Decode(data []byte, i interface{}) error {
	return json.Unmarshal(data, i)
}

// ProtobufCodec proto buf codec implement
type ProtobufCodec struct {
}

// Encode
func (c *ProtobufCodec) Encode(v interface{}) ([]byte, error) {
	if m, ok := v.(proto.Marshaler); ok {
		return m.Marshal()
	}

	if m, ok := v.(pb.Message); ok {
		return pb.Marshal(m)
	}

	return nil, fmt.Errorf("%T is not a proto.Marshaler", i)
}

// Decode
func (c *ProtobufCodec) Decode(data []byte, v interface{}) error {
	if m, ok := v.(proto.Unmarshaler); ok {
		return m.Unmarshal(data)
	}

	if m, ok := v.(pb.Message); ok {
		return pb.Unmarshal(data, m)
	}

	return fmt.Errorf("%T is not a proto.Unmarshaler", i)
}

// MessagePackCodec
type MessagePackCodec struct {
}

func (c *MessagePackCodec) Encode(i interface{}) ([]byte, error) {
	return nil, nil
}

func (c *MessagePackCodec) Decode(data []byte, i interface{}) error {
	return nil
}
