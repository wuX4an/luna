//go:build js && wasm

package console

import (
	"syscall/js"

	lua "github.com/yuin/gopher-lua"
)

// ---------------------------
// Loader del módulo
// ---------------------------
func Loader(L *lua.LState) *lua.LTable {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"log":   Log,
		"warn":  Warn,
		"error": Error,
		"info":  Info,
	})
	return mod
}

// ---------------------------
// Funciones
// ---------------------------
func Log(L *lua.LState) int {
	args := []any{}
	for i := 1; i <= L.GetTop(); i++ {
		args = append(args, L.ToString(i))
	}
	js.Global().Get("console").Call("log", args...)
	return 0
}

func Warn(L *lua.LState) int {
	args := []any{}
	for i := 1; i <= L.GetTop(); i++ {
		args = append(args, L.ToString(i))
	}
	js.Global().Get("console").Call("warn", args...)
	return 0
}

func Error(L *lua.LState) int {
	args := []any{}
	for i := 1; i <= L.GetTop(); i++ {
		args = append(args, L.ToString(i))
	}
	js.Global().Get("console").Call("error", args...)
	return 0
}

func Info(L *lua.LState) int {
	args := []any{}
	for i := 1; i <= L.GetTop(); i++ {
		args = append(args, L.ToString(i))
	}
	js.Global().Get("console").Call("info", args...)
	return 0
}
