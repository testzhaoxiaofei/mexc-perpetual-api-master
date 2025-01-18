package utils

import "encoding/base64"

// base64字符串转字节数组
func Base64ToBytes(base64Str string) ([]byte, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}

	return decodedBytes, nil
}

// 字节数组转base64字符串
func BytesToBase64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}
