package provider

import (
	"crypto/aes"
	"crypto/cipher"
)

type AesCtr struct {
	key    []byte
	iv     []byte
	stream cipher.Stream
}

func NewAesCtr(key []byte) *AesCtr {
	instance := &AesCtr{}
	instance.init(key)
	return instance
}

func (ins *AesCtr) init(key []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		println("EncryptAesCtr init NewCipher error")
	}

	blockSize := block.BlockSize()
	iv := []byte(key[:blockSize])

	stream := cipher.NewCTR(block, iv)

	ins.key = key
	ins.iv = iv
	ins.stream = stream
}

func (ins *AesCtr) Encrypt(bytes []byte) []byte {
	cipherBytes := make([]byte, len(bytes))
	ins.stream.XORKeyStream(cipherBytes, bytes)
	return cipherBytes
}

func (ins *AesCtr) Decrypt(cipherBytes []byte) []byte {
	bytes := make([]byte, len(cipherBytes))
	ins.stream.XORKeyStream(bytes, cipherBytes)
	return bytes
}
