//go:build !js

package http

import lua "github.com/yuin/gopher-lua"

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"request": Request,
		"server":  ServerProposal,
	})
	L.Push(mod)
	return 1
}
