package luavm

import (
	"luna/std"

	lua "github.com/yuin/gopher-lua"
)

func NewLuaVM() *lua.LState {
	L := lua.NewState(lua.Options{
		SkipOpenLibs: true,
	})
	// Registrar cierre en quien use esta función
	// Pero mejor que la llame quien la use (defer L.Close())

	lua.OpenBase(L)
	lua.OpenPackage(L)
	lua.OpenString(L)

	std.RegisterAll(L)

	return L
}
