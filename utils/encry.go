package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
)

var log *zap.SugaredLogger

func init() {
	log = NewLog()
}

func Encrypt(password, key string) (string, error) {
	// 随机64比特字符
	ranString64 := make([]byte, 64)
	io.ReadFull(rand.Reader, ranString64)

	b := buffer.Buffer{}
	b.Write(ranString64)
	b.Write([]byte(password))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("Encrypt: %v", err)
		return "", err
	}
	iv := make([]byte, aes.BlockSize)
	io.ReadFull(rand.Reader, iv)

	ciphertext := make([]byte, b.Len())
	stream := cipher.NewCFBEncrypter(block, iv) // Encrypt
	stream.XORKeyStream(ciphertext, b.Bytes())  // 参数：加密前/后的文本区
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(encryText, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("Encrypt: %v", err)
		return ""
	}
	// Base64 解密
	bytes, err := base64.StdEncoding.DecodeString(encryText)
	if err != nil {
		log.Fatalf("Decrypt: %v", err)
	}
	iv := make([]byte, aes.BlockSize)

	decrypted := make([]byte, len(bytes))
	stream := cipher.NewCFBDecrypter(block, iv) // Decrypt
	stream.XORKeyStream(decrypted, bytes)
	return string(decrypted[64:]) // 获取原始字符串
}
