//go:build !js

package sqlite

import lua "github.com/yuin/gopher-lua"

func luaDbExec(L *lua.LState) int {
	ud := checkDb(L)
	query := L.CheckString(2)
	_, err := ud.db.Exec(query)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LTrue)
	return 1
}
