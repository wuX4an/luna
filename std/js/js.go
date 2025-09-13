//go:build js && wasm
// +build js,wasm

package js

import (
	"syscall/js"

	lua "github.com/yuin/gopher-lua"
)

func Log(L *lua.LState) int {
	msg := L.CheckString(1)
	js.Global().Get("console").Call("log", msg)
	return 0
}

func Get(L *lua.LState) int {
	name := L.CheckString(1)
	val := js.Global().Get(name)
	switch val.Type() {
	case js.TypeString:
		L.Push(lua.LString(val.String()))
	case js.TypeNumber:
		L.Push(lua.LNumber(val.Float()))
	case js.TypeBoolean:
		L.Push(lua.LBool(val.Bool()))
	default:
		L.Push(lua.LNil)
	}
	return 1
}

func Set(L *lua.LState) int {
	name := L.CheckString(1)
	val := L.CheckAny(2)
	jsVal := js.Undefined()
	switch val.Type() {
	case lua.LTString:
		jsVal = js.ValueOf(val.String())
	case lua.LTNumber:
		jsVal = js.ValueOf(float64(val.(lua.LNumber)))
	case lua.LTBool:
		jsVal = js.ValueOf(bool(val.(lua.LBool)))
	}
	js.Global().Set(name, jsVal)
	return 0
}

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"log": Log,
		"get": Get,
		"set": Set,
	})
	L.Push(mod)
	return 1
}
