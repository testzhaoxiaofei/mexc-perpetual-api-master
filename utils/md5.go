package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// CalculateMD5 计算给定字符串的MD5值
func CalculateMD5(input string) string {
	// 创建一个MD5哈希对象
	hash := md5.New()

	// 将字符串写入哈希对象
	hash.Write([]byte(input))

	// 计算MD5值
	hashInBytes := hash.Sum(nil)

	// 将字节切片转换为16进制字符串
	md5String := hex.EncodeToString(hashInBytes)

	return md5String
}
