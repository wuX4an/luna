//go:build js && wasm
// +build js,wasm

package clipboard

import (
	"syscall/js"

	lua "github.com/yuin/gopher-lua"
)

// Loader devuelve el módulo clipboard listo para usar con await
func Loader(L *lua.LState) *lua.LTable {
	mod := L.NewTable()
	mod.RawSetString("read", L.NewFunction(Read))
	mod.RawSetString("write", L.NewFunction(Write))
	mod.RawSetString("clear", L.NewFunction(Clear))
	return mod
}

func Read(L *lua.LState) int {
	callback := L.CheckFunction(1) // recibe callback
	promise := js.Global().Get("navigator").Get("clipboard").Call("readText")

	var thenFunc js.Func
	thenFunc = js.FuncOf(func(this js.Value, args []js.Value) any {
		text := args[0].String()
		L.CallByParam(lua.P{
			Fn:      callback,
			NRet:    0,
			Protect: true,
		}, lua.LString(text))
		thenFunc.Release()
		return nil
	})

	promise.Call("then", thenFunc)
	return 0
}

// Write(text): suspende la coroutine hasta escribir en el clipboard
func Write(L *lua.LState) int {
	text := L.CheckString(1)
	resultChan := make(chan struct{}, 1)

	promise := js.Global().Get("navigator").Get("clipboard").Call("writeText", text)

	var thenFunc js.Func
	thenFunc = js.FuncOf(func(this js.Value, args []js.Value) any {
		resultChan <- struct{}{}
		thenFunc.Release()
		return nil
	})

	promise.Call("then", thenFunc)

	// suspender la coroutine hasta completarse
	<-resultChan
	return 0
}

// Clear(): limpia el portapapeles
func Clear(L *lua.LState) int {
	_ = js.Global().Get("navigator").Get("clipboard").Call("writeText", "")
	L.Push(lua.LBool(true))
	return 1
}
