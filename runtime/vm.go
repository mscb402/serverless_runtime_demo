package runtime

import (
	"fmt"
	"github.com/second-state/WasmEdge-go/wasmedge"
	"github.com/spf13/cast"
	"log"
	"os"
	"serverless_runtime_demo/lib"
)

func Run(code []byte, funcname string, args []string) (interface{}, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	// 初始化vm
	cfg, store := vmInit()
	vm := wasmedge.NewVMWithConfigAndStore(cfg, store)
	defer vm.Release()
	defer cfg.Release()
	defer store.Release()

	// 加载wasi
	var wasi = vm.GetImportObject(wasmedge.WASI)
	wasi.InitWasi(
		args,            // The args
		os.Environ(),    // The envs
		[]string{".:."}, // The mapping directories
	)

	// 加载模块代码
	err := vm.LoadWasmBuffer(code)
	if err != nil {
		return nil, err
	}

	// 验证模块代码
	err = vm.Validate()
	if err != nil {
		return nil, err
	}

	// 实例化
	err = vm.Instantiate()
	if err != nil {
		return nil, err
	}

	// 获取函数列表
	funcNames, funcTypes := vm.GetFunctionList()
	if len(funcNames) == 0 {
		return nil, fmt.Errorf("functions is empty")
	}

	// 检查函数是否存在
	i := lib.CheckFuncNameExist(funcname, funcNames)
	if i == -1 {
		return nil, fmt.Errorf("function %s not found", funcname)
	}

	log.Println("funcNames:", funcNames, "funcTypes:", funcTypes[i].GetParameters())

	// 函数入参
	params := funcTypes[i].GetParameters()
	paramsLen := len(params)

	fmt.Println("funcTypes:", params)

	if paramsLen != len(args) {
		return nil, fmt.Errorf("function %s args length is not equal %d != %d", funcname, paramsLen, len(args))
	}

	var ret interface{}

	// 有参数无参数需要分开处理
	if len(args) > 0 {
		passArgs := castType(args, params)
		ret, err = vm.Execute(funcname, passArgs...)
		if err != nil {
			return nil, err
		}
	} else {
		ret, err = vm.Execute(funcname)
		if err != nil {
			return nil, err
		}

	}
	// 返回结果
	return ret, nil

}

func vmInit() (cfg *wasmedge.Configure, store *wasmedge.Store) {
	cfg = wasmedge.NewConfigure(wasmedge.WASI)
	store = wasmedge.NewStore()
	return
}

// 将参数类型转换为目标类型
func castType(args []string, params []wasmedge.ValType) []interface{} {
	passArgs := []interface{}{}
	for i, arg := range args {
		if params[i] == wasmedge.ValType_I32 {
			passArgs = append(passArgs, cast.ToInt32(arg))
		} else if params[i] == wasmedge.ValType_I64 {
			passArgs = append(passArgs, cast.ToInt64(arg))
		} else if params[i] == wasmedge.ValType_F32 {
			passArgs = append(passArgs, cast.ToFloat32(arg))
		} else if params[i] == wasmedge.ValType_F64 {
			passArgs = append(passArgs, cast.ToFloat64(arg))
		} else {
			passArgs = append(passArgs, arg)
		}
	}
	return passArgs
}
