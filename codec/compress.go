package codec

var (
	compressors = make(map[string]Compressor, 0)
)

// Compressor
type Compressor interface {
	// Zip
	Zip(v interface{}) ([]byte, error)

	// Unzip
	Unzip(data []byte, v interface{}) error
}

// RegisterCompressor
func RegisterCompressor(name string, compressor Compressor) {
	compressors[name] = compressor
}

// GzipCompressor
type GzipCompressor struct {}

func (c GzipCompressor) Zip(data []byte) ([]byte, error) {
	return nil, nil
}

func (c GzipCompressor) Unzip(data []byte) ([]byte, error) {
	return nil, nil
}

type RawDataCompressor struct {
}

func (c RawDataCompressor) Zip(data []byte) ([]byte, error) {
	return data, nil
}

func (c RawDataCompressor) Unzip(data []byte) ([]byte, error) {
	return data, nil
}

// SnappyCompressor implements snappy compressor
type SnappyCompressor struct{}

func (c *SnappyCompressor) Zip(data []byte) ([]byte, error) {
	return nil, nil
}

func (c *SnappyCompressor) Unzip(data []byte) ([]byte, error) {
	return nil, nil
}
