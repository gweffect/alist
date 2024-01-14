package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"path"
	"strings"
)

const encodeURL = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

type FileNameBase64 struct {
	encoder *base64.Encoding
}

func NewFileNameBase64() *FileNameBase64 {
	ins := &FileNameBase64{}
	ins.init()
	return ins
}

func (ins *FileNameBase64) init() {
	ins.encoder = base64.NewEncoding(encodeURL)
}

func (ins *FileNameBase64) Encrypt(src string) string {
	// key := []byte("wumansgygoaescbc")
	// buffer, err := aesEncrypt(key, []byte(src))
	// if err != nil {

	// }
	fileExt := path.Ext(src)
	fileName := path.Base(src)
	fileNameWithoutExt := strings.TrimSuffix(fileName, fileExt)
	return ins.encoder.EncodeToString([]byte(fileNameWithoutExt)) + fileExt
}

func (ins *FileNameBase64) Decrypt(src string) string {
	if src == "root" {
		return src
	}
	fileExt := path.Ext(src)
	fileName := path.Base(src)
	fileNameWithoutExt := strings.TrimSuffix(fileName, fileExt)
	b1, err := ins.encoder.DecodeString(fileNameWithoutExt)
	if err != nil {
		return src
	}
	name := string(b1) + fileExt
	for _, e := range name {
		if e == 65533 {
			return src
		}
	}
	if strings.IndexAny(name, "\\x") > -1 {
		return src
	}
	// key := []byte("wumansgygoaescbc")
	// buffer, err := aesDecrypt(key, b1)

	return name
}

func aesEncrypt(key []byte, plainText []byte) ([]byte, error) {
	// 1. 使用密钥初始化
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	// 2. 对明文做 Padding 处理
	plainText = paddingPKCS7(plainText, blockSize)
	// if err != nil {
	//   return nil, err
	// }

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	cipherText := make([]byte, blockSize+sha256.Size+len(plainText))
	iv := cipherText[:blockSize]
	mac := cipherText[blockSize : blockSize+sha256.Size]
	payload := cipherText[blockSize+sha256.Size:]
	// 3. 初始化 IV
	if _, err = rand.Read(iv); err != nil {
		return nil, err
	}

	// 4. 加密明文
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(payload, plainText)

	// we use Encrypt-then-MAC
	// https://crypto.stackexchange.com/questions/202/should-we-mac-then-encrypt-or-encrypt-then-mac
	// 5. 计算 HMAC
	hash := hmac.New(sha256.New, key)
	hash.Write(payload)
	copy(mac, hash.Sum(nil))

	return cipherText, nil
}

func aesDecrypt(key []byte, cipherText []byte) ([]byte, error) {
	// 1. 使用密钥初始化
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	// 2. 验证长度是否过短
	if len(cipherText) <= blockSize+sha256.Size {
		return nil, errors.New("ciphertext too short")
	}

	iv := cipherText[:blockSize]
	mac := cipherText[blockSize : blockSize+sha256.Size]
	cipherText = cipherText[blockSize+sha256.Size:]

	// 2. 验证实际密文长度是否合法
	if len(cipherText)%blockSize != 0 {
		return nil, errors.New("ciphertext is not block-aligned, maybe corrupted")
	}

	hash := hmac.New(sha256.New, key)
	hash.Write(cipherText)
	// 3. 验证 HMAC
	if !hmac.Equal(hash.Sum(nil), mac) {
		return nil, errors.New("hmac failure, message corrupted")
	}

	plainText := make([]byte, len(cipherText))
	mode := cipher.NewCBCDecrypter(block, iv)
	// 4. 解密
	mode.CryptBlocks(plainText, cipherText)

	// 5. 去除 Padding
	plainText = unPaddingPKCS7(plainText)
	// if err != nil {
	//   return nil, err
	// }
	return plainText, nil
}

func paddingPKCS7(plaintext []byte, blockSize int) []byte {
	paddingSize := blockSize - len(plaintext)%blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(plaintext, paddingText...)
}

// PKCS7 反填充
func unPaddingPKCS7(s []byte) []byte {
	length := len(s)
	if length == 0 {
		return s
	}
	unPadding := int(s[length-1])
	return s[:(length - unPadding)]
}
