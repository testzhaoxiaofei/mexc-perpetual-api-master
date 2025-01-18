package utils

import (
	"crypto/rand"
)

// GenerateRandomBytes 生成指定长度的随机字节切片
func GenerateRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	return randomBytes, nil
}
