package jsext

import (
	"fmt"

	"github.com/dop251/goja"
)

var builtinNatives = map[string]func(*goja.Runtime) func(call goja.FunctionCall) goja.Value{
	"print": consolePrint,
}

func consolePrint(vm *goja.Runtime) func(call goja.FunctionCall) goja.Value {
	return func(call goja.FunctionCall) goja.Value {
		printArgs := make([]interface{}, len(call.Arguments))
		for i, arg := range call.Arguments {
			printArgs[i] = arg
		}
		fmt.Println(printArgs...)
		return vm.ToValue(goja.Undefined())
	}
}

func RegisterBuiltinNatives(vm *goja.Runtime) {
	for name, function := range builtinNatives {
		vm.Set(name, function(vm))
	}
}
