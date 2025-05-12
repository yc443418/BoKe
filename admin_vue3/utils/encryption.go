// 加密工具类
// @author chen

package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptionMd5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
