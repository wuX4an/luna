//go:build js && wasm

package stub

import (
	"fmt"
	"syscall/js"

	lua "github.com/yuin/gopher-lua"
)

func NewStub(modName string) lua.LGFunction {
	return func(L *lua.LState) int {
		msg := fmt.Sprintf("Warning: module %s is not implemented on JS/WASM", modName)

		// Mostrar en consola JS (browser, Node, Bun)
		js.Global().Get("console").Call("warn", msg)

		// Retornar tabla vacía
		L.Push(L.NewTable())
		return 1
	}
}
