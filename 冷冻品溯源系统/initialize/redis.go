package initialize

import (
	"Traceability/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

var Ctx = context.Background()

func InitRedis() {
	var err error
	db, err := strconv.Atoi(global.CONFIG.Redis.DB)
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", global.CONFIG.Redis.Host, global.CONFIG.Redis.Port),
		Password: global.CONFIG.Redis.Password, // no password set
		DB:       db,                           // use default DB
	})
	fmt.Println("redis", global.Redis)
	// 测试连接
	_, err = global.Redis.Ping(Ctx).Result()
	if err != nil {
		log.Fatalln("连接redis失败：" + err.Error())
		return
	}
	fmt.Print("Redis连接成功")
}

func CloseRedis() {
	err := global.Redis.Close()
	if err != nil {
		log.Fatalln("关闭redis失败：" + err.Error())
		return
	}
	log.Fatalln("Redis连接已关闭")
}
