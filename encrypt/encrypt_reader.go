package encrypt

import (
	"io"

	"github.com/gweffectx/safedav/encrypt/provider"
)

type EncryptReader struct {
	reader   io.ReadCloser
	provider *provider.AesCtr
}

func NewEncryptReader(reader io.ReadCloser, key []byte) *EncryptReader {
	ins := &EncryptReader{}
	ins.init(reader, key)
	return ins
}

func (ins *EncryptReader) init(reader io.ReadCloser, key []byte) {
	provider := provider.NewAesCtr(key)
	ins.provider = provider
	ins.reader = reader
}

func (ins *EncryptReader) Read(buffer []byte) (n int, err error) {
	_buffer := make([]byte, len(buffer))
	readLength, err := ins.reader.Read(_buffer)
	if err != nil {
		return readLength, err
	}
	if readLength == 0 {
		return 0, nil
	}
	data := ins.provider.Encrypt(_buffer[:readLength])
	dataLength := len(data)
	for i := 0; i < dataLength; i++ {
		buffer[i] = data[i]
	}
	return dataLength, nil
}

func (ins *EncryptReader) Close() error {
	return ins.reader.Close()
}
