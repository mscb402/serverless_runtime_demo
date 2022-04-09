package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"serverless_runtime_demo/runtime"
	"strings"
)

func RunFunc(ctx *gin.Context) {
	funcHash := ctx.Param("hash")
	funcName := ctx.Query("func")
	arg := ctx.Query("arg")
	args := []string{}
	if arg != "" {
		args = strings.Split(arg, ",")
	}

	code, err := LoaderInstance.Load(funcHash)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	if len(code) == 0 {
		ctx.JSON(200, gin.H{
			"code": -1,
			"msg":  "code is empty",
		})
		return
	}
	ret, err := runtime.Run(code, funcName, args)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	fmt.Println(ret)
	ctx.JSON(200, gin.H{
		"code": 0,
		"data": ret,
	})
}
