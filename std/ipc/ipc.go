package ipc

import (
	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"server": Server,
		"client": Client,
	})
	L.Push(mod)
	return 1
}
