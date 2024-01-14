package encrypt

import (
	"io"

	"github.com/gweffectx/safedav/encrypt/provider"
)

type DecryptReader struct {
	reader   io.ReadCloser
	provider *provider.AesCtr
}

func NewDecryptReader(reader io.ReadCloser, key []byte) *DecryptReader {
	ins := &DecryptReader{}
	ins.init(reader, key)
	return ins
}

func (ins *DecryptReader) init(reader io.ReadCloser, key []byte) {
	provider := provider.NewAesCtr(key)
	ins.provider = provider
	ins.reader = reader
}

func (ins *DecryptReader) SetOffset(offset int64) {
	ins.provider.SetOffset(offset)
}

func (ins *DecryptReader) Read(buffer []byte) (n int, err error) {
	_buffer := make([]byte, len(buffer))
	readLength, err := ins.reader.Read(_buffer)
	if err != nil {
		//return readLength, err
	}
	if readLength == 0 {
		return 0, err
	}
	data := ins.provider.Decrypt(_buffer[:readLength])
	dataLength := len(data)
	for i := 0; i < dataLength; i++ {
		buffer[i] = data[i]
	}
	return readLength, err
}

func (ins *DecryptReader) Close() error {
	return ins.reader.Close()
}
