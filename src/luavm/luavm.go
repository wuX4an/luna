package luavm

import (
	"luna/std"

	lua "github.com/yuin/gopher-lua"
)

func NewLuaVM() *lua.LState {
	L := lua.NewState()

	std.RegisterAll(L)

	return L
}
