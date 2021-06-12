package scan

import (
	"io"
	"log"

	"github.com/Snawoot/trusearch/def"
	"github.com/Snawoot/trusearch/jsext"

	"github.com/dop251/goja"
)

func Scan(scanner def.RecordScanner, funcCode string) int {
	vm := goja.New()
	jsext.RegisterBuiltinNatives(vm)

	jsFunValue, err := vm.RunString(funcCode)
	if err != nil {
		log.Printf("Function fails to compile: %v", err)
		return 5
	}

	jsFun, ok := goja.AssertFunction(jsFunValue)
	if !ok {
		log.Printf("JS code doesn't evaluates to a callable function")
		return 6
	}

	for {
		rec, err := scanner.Scan()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Got error on input scan: %v", err)
			return 3
		}

		_, err = jsFun(goja.Undefined(), vm.ToValue(*rec))
		if err != nil {
			log.Printf("Function invocation failed: %v", err)
		}
	}

	return 0
}
