package sqlite

import (
	"database/sql"

	lua "github.com/yuin/gopher-lua"
)

type dbUserData struct {
	db *sql.DB
}

func registerRowsType(L *lua.LState) {
	mt := L.NewTypeMetatable("sqlite_rows")
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"next":  luaRowsNext,
		"close": luaRowsClose,
	}))
}

func checkDb(L *lua.LState) *dbUserData {
	ud := L.CheckUserData(1)
	db, ok := ud.Value.(*dbUserData)
	if !ok {
		L.ArgError(1, "expected sqlite db")
	}
	return db
}

func pushDB(L *lua.LState, db *sql.DB) int {
	ud := L.NewUserData()
	ud.Value = &dbUserData{db: db}
	L.SetMetatable(ud, L.GetTypeMetatable("sqlite_db"))
	L.Push(ud)
	return 1
}

func Loader(L *lua.LState) int {
	registerRowsType(L)
	registerTxType(L)

	mt := L.NewTypeMetatable("sqlite_db")
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"exec":        luaDbExec,
		"query":       luaDbQuery,
		"transaction": luaDbTransaction,
		"close": func(L *lua.LState) int {
			ud := checkDb(L)
			ud.db.Close()
			return 0
		},
	}))

	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"open":      luaOpen,
		"in_memory": luaInMemory,
	})
	L.Push(mod)
	return 1
}
