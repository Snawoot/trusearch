package jsutil

import (
	"errors"

	"github.com/dop251/goja"
)

var (
	ErrUndefined    = errors.New("value is not defined")
	ErrNotAFunction = errors.New("value is not a function")
)

func FnInvoke(vm *goja.Runtime, name string, args ...goja.Value) (goja.Value, error) {
	val := vm.Get(name)
	if val == nil {
		return nil, ErrUndefined
	}
	fn, ok := goja.AssertFunction(val)
	if !ok {
		return nil, ErrNotAFunction
	}

	return fn(goja.Undefined(), args...)
}
