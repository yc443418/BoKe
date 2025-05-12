package initialize

import (
	"Traceability/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func TestMustConfig() {
	var config string
	flag.StringVar(&config, "config", "", "指定文件夹路径")
	flag.Parse()
	if config == "" {
		if configEnv := os.Getenv("CONFIG"); configEnv != "" {
			switch gin.Mode() {
			case gin.DebugMode:
				config = "./config/config.yaml"
			case gin.TestMode:
				config = "./config/config.test.yaml"
			case gin.ReleaseMode:
				config = "./config/config.release.yaml"
			}
			fmt.Printf("您正在使用gin模式的%s环境名称, 配置文件的路径为%s\n", gin.Mode(), config)
		} else {
			config = configEnv
			fmt.Printf("您正在使用环境变量, 配置文件的路径为%s\n", configEnv)
		}
	} else {
		fmt.Printf("您正在使用命令行的-c参数传递的值, 配置文件的路径为%s\n", config)
	}
	// 绑定到结构体中
	v := viper.New()
	v.SetConfigName("../config/config_test")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(&global.CONFIG)
	fmt.Print("jwt", global.CONFIG.Jwt)
	if err != nil {
		panic(err)
	}
	v.WatchConfig()

	// 配置热加载功能
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生变更: ", in.Name)
		err = v.Unmarshal(&global.CONFIG)
		if err != nil {
			panic(err)
		}
		fmt.Println(global.CONFIG)
	})
	select {}
}
