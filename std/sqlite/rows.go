//go:build !js

package sqlite

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func luaRowsNext(L *lua.LState) int {
	ud := L.CheckUserData(1)
	sd, ok := ud.Value.(*stmtUserData)
	if !ok {
		L.ArgError(1, "expected sqlite rows")
		return 0
	}
	if !sd.rows.Next() {
		L.Push(lua.LNil)
		return 1
	}

	cols, err := sd.rows.Columns()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	values := make([]interface{}, len(cols))
	valuePtrs := make([]interface{}, len(cols))
	for i := range values {
		valuePtrs[i] = &values[i]
	}
	if err := sd.rows.Scan(valuePtrs...); err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	row := L.NewTable()
	for i, col := range cols {
		val := values[i]
		switch v := val.(type) {
		case int64:
			row.RawSetString(col, lua.LNumber(v))
		case float64:
			row.RawSetString(col, lua.LNumber(v))
		case bool:
			row.RawSetString(col, lua.LBool(v))
		case []byte:
			row.RawSetString(col, lua.LString(string(v)))
		case string:
			row.RawSetString(col, lua.LString(v))
		case nil:
			row.RawSetString(col, lua.LNil)
		default:
			row.RawSetString(col, lua.LString(fmt.Sprintf("%v", v)))
		}
	}
	L.Push(row)
	return 1
}

func luaRowsClose(L *lua.LState) int {
	ud := L.CheckUserData(1)
	sd, ok := ud.Value.(*stmtUserData)
	if !ok {
		L.ArgError(1, "expected sqlite rows")
		return 0
	}
	if err := sd.rows.Close(); err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LTrue)
	return 1
}

func luaRowsIter(L *lua.LState) int {
	ud := L.CheckUserData(1)
	sd, ok := ud.Value.(*stmtUserData)
	if !ok {
		L.ArgError(1, "expected sqlite rows")
		return 0
	}

	// Función iteradora
	iter := L.NewFunction(func(L *lua.LState) int {
		if !sd.rows.Next() {
			L.Push(lua.LNil)
			return 1
		}

		cols, err := sd.rows.Columns()
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		values := make([]interface{}, len(cols))
		valuePtrs := make([]interface{}, len(cols))
		for i := range values {
			valuePtrs[i] = &values[i]
		}
		if err := sd.rows.Scan(valuePtrs...); err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		row := L.NewTable()
		for i, col := range cols {
			val := values[i]
			switch v := val.(type) {
			case int64:
				row.RawSetString(col, lua.LNumber(v))
			case float64:
				row.RawSetString(col, lua.LNumber(v))
			case bool:
				row.RawSetString(col, lua.LBool(v))
			case []byte:
				row.RawSetString(col, lua.LString(string(v)))
			case string:
				row.RawSetString(col, lua.LString(v))
			case nil:
				row.RawSetString(col, lua.LNil)
			default:
				row.RawSetString(col, lua.LString(fmt.Sprintf("%v", v)))
			}
		}
		L.Push(row)
		return 1
	})

	L.Push(iter)
	return 1
}
