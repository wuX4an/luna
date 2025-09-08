package json

import (
	lua "github.com/yuin/gopher-lua"
)

// Loader registers the module in Lua
func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"encode": Encode,
		"decode": Decode,
	})
	L.Push(mod)
	return 1
}
