package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	temstr := h.Sum(nil)
	return hex.EncodeToString(temstr)
}

// 大写
func Md5Decode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

func MakePassword(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}
