package encrypt

import (
	"io"

	"github.com/gweffectx/safedav/encrypt/provider"
)

type EncryptWriter struct {
	writer   io.Writer
	provider *provider.AesCtr
}

func NewEncryptWriter(writer io.Writer, key []byte) *EncryptWriter {
	ins := &EncryptWriter{}
	ins.init(writer, key)
	return ins
}

func (ins *EncryptWriter) init(writer io.Writer, key []byte) {
	provider := provider.NewAesCtr(key)
	ins.provider = provider
	ins.writer = writer
}

func (ins *EncryptWriter) Write(data []byte) (n int, err error) {
	buffer := ins.provider.Encrypt(data)
	return ins.writer.Write(buffer)
}
