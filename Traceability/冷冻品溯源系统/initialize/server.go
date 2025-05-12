package initialize

import (
	"fmt"
	"gorm.io/gorm"
	_ "net/http/pprof"

	"Traceability/global"
	"Traceability/router"
	"github.com/gin-gonic/gin"
)

func MustRunWebServer(db *gorm.DB) {
	engine := gin.Default()

	go router.InitRouter(engine, db)
	address := fmt.Sprintf(":%d", global.CONFIG.Server.Port)
	fmt.Println("启动服务器，监听端口：", address)
	if err := engine.Run(address); err != nil {
		panic(err)
	}
}
