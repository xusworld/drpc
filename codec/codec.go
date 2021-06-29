// codec
package codec

var (
	codecs = make(map[string]Codec, 0)
)

// Codec Interface
type Codec interface {
	// Encode
	Encode(req interface{}) ([]byte,error)

	// Decode
	Decode(req, rsp interface{}) error
}

// RegisterCodec
func RegisterCodec(name string, codec Codec) Codec {
	codecs[name] = codec
	return nil
}

// GetCodec
func GetCodec(name string) Codec {
	return codecs[name]
}
