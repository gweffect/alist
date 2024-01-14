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

func (ins *AesCtr) SetOffset(offset int64) {
	block, err := aes.NewCipher(ins.key)
	if err != nil {
		println("EncryptAesCtr init NewCipher error")
	}

	blockSize := int64(block.BlockSize())
	iv := []byte(ins.key[:blockSize])

	// // 计算偏移量对应的块索引和字节索引
	// blockIndex := offset / blockSize
	// //byteIndex := offset % blockSize
	// // 计算偏移量对应的计数器值
	// counter := blockIndex
	// // 将计数器值转换为字节序列
	// counterBytes := make([]byte, blockSize)
	// binary.BigEndian.PutUint64(counterBytes, uint64(counter))

	// stream := cipher.NewCTR(block, counterBytes)

	// ins.iv = counterBytes
	// ins.stream = stream

	chunks := uint64(int(offset) / block.BlockSize())

	// iv += offset
	var c uint16
	for i := len(iv[:]) - 1; i >= 0; i-- {
		c = uint16(iv[i]) + uint16(chunks&0xFF) + c
		iv[i], c, chunks = byte(c), c>>8, chunks>>8
	}

	// Reinitialize cipher
	stream := cipher.NewCTR(block, iv[:])

	ins.iv = iv
	ins.stream = stream
}

func (ins *AesCtr) IncrementIV1(offset int64) {
	block, err := aes.NewCipher(ins.key)
	if err != nil {
		println("EncryptAesCtr init NewCipher error")
	}

	blockSize := block.BlockSize()
	iv := []byte(ins.key[:blockSize])
	increment := (int(offset) / blockSize)

	const MAX_UINT32 = 0xffffffff
	var incrementBig = ^(increment / MAX_UINT32)
	var incrementLittle = increment%MAX_UINT32 - incrementBig
	// split the 128bits IV in 4 numbers, 32bits each
	overflow := 0
	for idx := 0; idx < 4; idx++ {
		num := int(iv[12-idx*4]) | (int(iv[12-idx*4+1]) << 8) | (int(iv[12-idx*4+2]) << 16) | (int(iv[12-idx*4+3]) << 24)
		inc := overflow
		if idx == 0 {
			inc += incrementLittle
		}
		if idx == 1 {
			inc += incrementBig
		}
		num += inc
		numBig := num / MAX_UINT32
		numLittle := num%MAX_UINT32 - numBig
		overflow = numBig
		iv[12-idx*4] = byte(numLittle)
		iv[12-idx*4+1] = byte(numLittle >> 8)
		iv[12-idx*4+2] = byte(numLittle >> 16)
		iv[12-idx*4+3] = byte(numLittle >> 24)
	}

	stream := cipher.NewCTR(block, iv[:])

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
