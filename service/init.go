package service

import (
	"github.com/spf13/viper"
	"serverless_runtime_demo/code_loader"
)

var LoaderInstance code_loader.Loader

func Init(cfg *viper.Viper) {
	loaderType := cfg.GetString("code_loader")

	//init loader
	loader, err := code_loader.NewLoader(loaderType, cfg)
	if err != nil {
		panic(err)
	}
	if loader == nil {
		panic("init loader failed")
	}
	LoaderInstance = loader

}
