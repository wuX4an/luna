package sqlite

import (
	"database/sql"

	lua "github.com/yuin/gopher-lua"
	_ "modernc.org/sqlite"
)

func luaOpen(L *lua.LState) int {
	filename := L.CheckString(1)
	db, err := sql.Open("sqlite", filename)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	return pushDB(L, db)
}

func luaInMemory(L *lua.LState) int {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	return pushDB(L, db)
}
