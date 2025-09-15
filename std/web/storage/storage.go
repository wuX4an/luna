//go:build js && wasm

package storage

import (
	"syscall/js"

	lua "github.com/yuin/gopher-lua"
)

// ---------------------------
// Loader del módulo
// ---------------------------

func Loader(L *lua.LState) *lua.LTable {
	mod := L.NewTable()
	mod.RawSetString("localStorage", exportLocalStorage(L))
	mod.RawSetString("sessionStorage", exportSessionStorage(L))
	return mod
}

// ---------------------------
// LocalStorage
// ---------------------------
func exportLocalStorage(L *lua.LState) *lua.LTable {
	return L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"setItem":    func(L *lua.LState) int { return setItem(L, js.Global().Get("localStorage")) },
		"getItem":    func(L *lua.LState) int { return getItem(L, js.Global().Get("localStorage")) },
		"removeItem": func(L *lua.LState) int { return removeItem(L, js.Global().Get("localStorage")) },
		"clear":      func(L *lua.LState) int { return clear(js.Global().Get("localStorage")) },
	})
}

// ---------------------------
// SessionStorage
// ---------------------------
func exportSessionStorage(L *lua.LState) *lua.LTable {
	return L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"setItem":    func(L *lua.LState) int { return setItem(L, js.Global().Get("sessionStorage")) },
		"getItem":    func(L *lua.LState) int { return getItem(L, js.Global().Get("sessionStorage")) },
		"removeItem": func(L *lua.LState) int { return removeItem(L, js.Global().Get("sessionStorage")) },
		"clear":      func(L *lua.LState) int { return clear(js.Global().Get("sessionStorage")) },
	})
}

// ---------------------------
// Helpers
// ---------------------------
func setItem(L *lua.LState, store js.Value) int {
	key := L.CheckString(1)
	value := L.CheckString(2)
	store.Call("setItem", key, value)
	return 0
}

func getItem(L *lua.LState, store js.Value) int {
	key := L.CheckString(1)
	val := store.Call("getItem", key)
	if val.Type() == js.TypeNull || val.Type() == js.TypeUndefined {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(val.String()))
	}
	return 1
}

func removeItem(L *lua.LState, store js.Value) int {
	key := L.CheckString(1)
	store.Call("removeItem", key)
	return 0
}

func clear(store js.Value) int {
	store.Call("clear")
	return 0
}
