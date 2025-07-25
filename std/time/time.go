package time

import (
	lua "github.com/yuin/gopher-lua"
)

// Loader es el punto de entrada de la librería para require("time")
func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"now":    now,
		"sleep":  sleep,
		"format": format,
	})
	L.Push(mod)
	return 1
}
