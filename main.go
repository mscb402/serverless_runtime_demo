package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"serverless_runtime_demo/service"
)

func main() {
	//初始化
	cfg := configInit()
	service.Init(cfg)
	engine := gin.Default()

	// 注册路由
	route(engine)

	// 启动服务
	err := engine.Run(":8080")
	if err != nil {
		log.Panicln(err)
		return
	}
}
func route(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})
	engine.GET("/run/:hash", service.RunFunc)
}

func configInit() *viper.Viper {
	cfg := viper.New()
	cfg.SetConfigType("yaml")
	cfg.SetConfigFile("config.yaml")
	err := cfg.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return cfg
}
