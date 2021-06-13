package jsext

import (
	"fmt"
	"os"

	"github.com/Snawoot/bbcode"
	"github.com/dop251/goja"
)

var builtinNatives = map[string]func(*goja.Runtime) func(call goja.FunctionCall) goja.Value{
	"perror":       stderrPrint,
	"print":        stdoutPrint,
	"strip_bbcode": stripBBCode,
}

var bbCodeCompiler = bbcode.NewCompiler(true, true) // autoCloseTags, ignoreUnmatchedClosingTags

func init() {
	for _, tag := range []string{
		"i",
		"b",
		"u",
		"s",
		"size",
		"color",
		"table",
		"th",
		"tr",
		"td",
		"quote",
		"img",
		"url",
		"list",
		"hr",
		"code",
		"spoiler",
		"pre",
		"box",
		"nfo",
		"openline",
		"align",
		"font",
	} {
		bbCodeCompiler.SetTag(tag, func(node *bbcode.BBCodeNode) (*bbcode.HTMLTag, bool) {
			out := bbcode.NewHTMLTag("")
			return out, true
		})
	}
	for _, tag := range []string{"hr", "br"} {
		bbCodeCompiler.SetTag(tag, func(node *bbcode.BBCodeNode) (*bbcode.HTMLTag, bool) {
			out := bbcode.NewHTMLTag("\n")
			return out, true
		})
	}
}

func stdoutPrint(vm *goja.Runtime) func(call goja.FunctionCall) goja.Value {
	return func(call goja.FunctionCall) goja.Value {
		printArgs := make([]interface{}, len(call.Arguments))
		for i, arg := range call.Arguments {
			printArgs[i] = arg
		}
		fmt.Println(printArgs...)
		return vm.ToValue(goja.Undefined())
	}
}

func stderrPrint(vm *goja.Runtime) func(call goja.FunctionCall) goja.Value {
	return func(call goja.FunctionCall) goja.Value {
		printArgs := make([]interface{}, len(call.Arguments))
		for i, arg := range call.Arguments {
			printArgs[i] = arg
		}
		fmt.Fprintln(os.Stderr, printArgs...)
		return vm.ToValue(goja.Undefined())
	}
}

func stripBBCode(vm *goja.Runtime) func(call goja.FunctionCall) goja.Value {
	return func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) == 0 {
			return vm.ToValue(goja.Undefined())
		}
		return vm.ToValue(bbCodeCompiler.Compile(call.Arguments[0].String()))
	}
}

func RegisterBuiltinNatives(vm *goja.Runtime) {
	for name, function := range builtinNatives {
		vm.Set(name, function(vm))
	}
}
