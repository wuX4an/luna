//go:build !js

package stub

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func NewStub(modName string) lua.LGFunction {
	return func(L *lua.LState) int {
		// Mensaje con "WARNING" en amarillo
		msg := fmt.Sprintf("\033[33mWARNING\033[0m: module %s is only implemented on JS/WASM", modName)

		// Mostrar en Lua
		L.CallByParam(lua.P{
			Fn:      L.GetGlobal("print"),
			NRet:    0,
			Protect: true,
		}, lua.LString(msg))

		// Retornar tabla vacía
		L.Push(L.NewTable())
		return 1
	}
}
