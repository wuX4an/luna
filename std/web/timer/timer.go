//go:build js && wasm
// +build js,wasm

package timer

import (
	"syscall/js"

	lua "github.com/yuin/gopher-lua"
)

// ---------------------------
// Loader del módulo
// ---------------------------
func Loader(L *lua.LState) *lua.LTable {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"setTimeout":    SetTimeout,
		"setInterval":   SetInterval,
		"clearTimeout":  ClearTimeout,
		"clearInterval": ClearInterval,
		"sleep":         Sleep,
	})
	return mod
}

// ---------------------------
// Funciones
// ---------------------------
func SetTimeout(L *lua.LState) int {
	fn := L.CheckFunction(1)
	ms := L.CheckInt(2)
	id := js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, args []js.Value) any {
		L.CallByParam(lua.P{
			Fn:      fn,
			NRet:    0,
			Protect: true,
		})
		return nil
	}), ms)
	L.Push(lua.LNumber(id.Int()))
	return 1
}

func ClearTimeout(L *lua.LState) int {
	id := L.CheckInt(1)
	js.Global().Call("clearTimeout", id)
	return 0
}

func SetInterval(L *lua.LState) int {
	fn := L.CheckFunction(1)
	ms := L.CheckInt(2)
	id := js.Global().Call("setInterval", js.FuncOf(func(this js.Value, args []js.Value) any {
		L.CallByParam(lua.P{
			Fn:      fn,
			NRet:    0,
			Protect: true,
		})
		return nil
	}), ms)
	L.Push(lua.LNumber(id.Int()))
	return 1
}

func ClearInterval(L *lua.LState) int {
	id := L.CheckInt(1)
	js.Global().Call("clearInterval", id)
	return 0
}

// ---------------------------
// sleep(ms) -> awaitable
// ---------------------------
func Sleep(L *lua.LState) int {
	ms := L.CheckInt(1)

	// Creamos una promesa JS que se resuelve después de ms
	promise := js.Global().Get("Promise").New(js.FuncOf(func(this js.Value, args []js.Value) any {
		resolve := args[0]
		js.Global().Call("setTimeout", resolve, ms)
		return nil
	}))

	// Lo devolvemos como userdata para await
	ud := L.NewUserData()
	ud.Value = promise
	L.Push(ud)
	return 1
}
