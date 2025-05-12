package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

// RandSalt 用于生成随机盐值
func RandSalt() (string, error) {
	// 定义盐值长度（建议至少为 32 字节，以增强安全性）
	const saltLength = 32

	// 创建一个字节数组用于存储随机数
	saltBytes := make([]byte, saltLength)

	// 使用 crypto/rand 包生成安全的随机字节
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", errors.New("failed to generate random bytes")
	}

	// 将字节数组编码为 Base64 字符串
	salt := base64.StdEncoding.EncodeToString(saltBytes)

	return salt, nil
}
