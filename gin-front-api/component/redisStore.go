// redis存取验证码
// @author chen

package component

import (
	"gin-front-api/constant"
	"gin-front-api/core"
	"gin-front-api/global"
	"log"
	"time"
)

type RedisStore struct{}

// Set 存验证码
func (r RedisStore) Set(id string, value string) {
	key := constant.LoginCode + id
	err := core.RedisDb.Set(global.Ctx, key, value, time.Minute*5).Err()
	if err != nil {
		log.Panicln(err.Error())
	}
}

// Get 取验证码
func (r RedisStore) Get(id string, clear bool) string {
	key := constant.LoginCode + id
	value, err := core.RedisDb.Get(global.Ctx, key).Result()
	if err != nil {
		return ""
	}
	return value
}

// Verify 校验验证码
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
