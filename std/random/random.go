package random

import lua "github.com/yuin/gopher-lua"

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"choice": Choice,
		"take":   Take,
		"int":    Int,
		"float":  Float,
	})
	L.Push(mod)
	return 1
}
