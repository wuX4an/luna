package env

import lua "github.com/yuin/gopher-lua"

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"load": Load,
		"get":  Get,
		"set":  Set,
	})
	L.Push(mod)
	return 1
}
