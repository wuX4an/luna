//go:build js && wasm
// +build js,wasm

package notify

import (
	"syscall/js"

	lua "github.com/yuin/gopher-lua"
)

// Loader del módulo
func Loader(L *lua.LState) *lua.LTable {
	mod := L.NewTable()
	mod.RawSetString("requestPermission", L.NewFunction(RequestPermission))
	mod.RawSetString("show", L.NewFunction(Show))
	mod.RawSetString("permission", L.NewFunction(Permission))
	return mod
}

// requestPermission(callback)
// Se usa con await: await(notify.requestPermission)
func RequestPermission(L *lua.LState) int {
	resultChan := make(chan lua.LValue, 1)

	promise := js.Global().Get("Notification").Call("requestPermission")

	var thenFunc js.Func
	thenFunc = js.FuncOf(func(this js.Value, args []js.Value) any {
		resultChan <- lua.LString(args[0].String())
		thenFunc.Release()
		return nil
	})

	promise.Call("then", thenFunc)

	res := <-resultChan
	L.Push(res)
	return 1
}

// show(title, options)
// No es await, solo dispara la notificación
func Show(L *lua.LState) int {
	title := L.CheckString(1)
	opts := L.OptTable(2, L.NewTable())

	// Crear objeto JS para opciones
	jsOpts := js.Global().Get("Object").New()

	opts.ForEach(func(k, v lua.LValue) {
		key := k.String()
		switch val := v.(type) {
		case lua.LString:
			jsOpts.Set(key, string(val))
		case lua.LBool:
			jsOpts.Set(key, bool(val))
		}
	})

	// Crear la notificación
	js.Global().Get("Notification").New(title, jsOpts)
	return 0
}

// permission(): devuelve el estado actual ("granted", "denied", "default")
func Permission(L *lua.LState) int {
	perm := js.Global().Get("Notification").Get("permission").String()
	L.Push(lua.LString(perm))
	return 1
}
