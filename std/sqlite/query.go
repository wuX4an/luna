//go:build !js

package sqlite

import (
	"database/sql"

	lua "github.com/yuin/gopher-lua"
)

type stmtUserData struct {
	rows *sql.Rows
}

func luaDbQuery(L *lua.LState) int {
	ud := checkDb(L)
	sqlQuery := L.CheckString(2)

	rows, err := ud.db.Query(sqlQuery)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	udRows := L.NewUserData()
	udRows.Value = &stmtUserData{rows: rows}
	L.SetMetatable(udRows, L.GetTypeMetatable("sqlite_rows"))
	L.Push(udRows)
	return 1
}
