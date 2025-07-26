package sqlite

import (
	"database/sql"

	lua "github.com/yuin/gopher-lua"
)

// txUserData representa una transacción SQLite en Lua
type txUserData struct {
	tx *sql.Tx
}

// Registro del tipo sqlite_tx en Lua
func registerTxType(L *lua.LState) {
	mt := L.NewTypeMetatable("sqlite_tx")
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"exec":  luaTxExec,
		"query": luaTxQuery,
		"close": luaTxClose,
	}))
}

// Método exec para sqlite_tx: ejecuta SQL dentro de la transacción
func luaTxExec(L *lua.LState) int {
	ud := L.CheckUserData(1)
	txd, ok := ud.Value.(*txUserData)
	if !ok {
		L.ArgError(1, "sqlite transaction expected")
		return 0
	}
	sqlStr := L.CheckString(2)
	_, err := txd.tx.Exec(sqlStr)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LTrue)
	return 1
}

// Método query para sqlite_tx: consulta SQL dentro de la transacción
func luaTxQuery(L *lua.LState) int {
	ud := L.CheckUserData(1)
	txd, ok := ud.Value.(*txUserData)
	if !ok {
		L.ArgError(1, "sqlite transaction expected")
		return 0
	}
	sqlStr := L.CheckString(2)
	rows, err := txd.tx.Query(sqlStr)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	udRows := L.NewUserData()
	udRows.Value = &stmtUserData{rows: rows} // stmtUserData debe estar definido en tu código para manejar filas
	L.SetMetatable(udRows, L.GetTypeMetatable("sqlite_rows"))
	L.Push(udRows)
	return 1
}

// Método close para sqlite_tx: rollback manual (opcional)
func luaTxClose(L *lua.LState) int {
	ud := L.CheckUserData(1)
	txd, ok := ud.Value.(*txUserData)
	if !ok {
		L.ArgError(1, "sqlite transaction expected")
		return 0
	}
	err := txd.tx.Rollback()
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LTrue)
	return 1
}

// Función para iniciar la transacción y ejecutar el callback Lua
func luaDbTransaction(L *lua.LState) int {
	ud := checkDb(L)         // Obtenemos la base de datos
	fn := L.CheckFunction(2) // Callback Lua que recibe la transacción

	tx, err := ud.db.Begin()
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	// Creamos userdata para la transacción
	txUd := L.NewUserData()
	txUd.Value = &txUserData{tx: tx}
	L.SetMetatable(txUd, L.GetTypeMetatable("sqlite_tx"))

	// Ejecutamos callback Lua pasándole la transacción
	err = L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    0,
		Protect: true,
	}, txUd)

	if err != nil {
		// Error en callback -> rollback
		_ = tx.Rollback()
		L.Push(lua.LFalse)
		L.Push(lua.LString("Transaction callback error: " + err.Error()))
		return 2
	}

	// Commit de la transacción
	if err := tx.Commit(); err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString("Commit error: " + err.Error()))
		return 2
	}

	// Todo OK
	L.Push(lua.LTrue)
	return 1
}
