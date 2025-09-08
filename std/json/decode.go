package json

import (
	"encoding/json"

	lua "github.com/yuin/gopher-lua"
)

// Decode: converts JSON string to Lua table
func Decode(L *lua.LState) int {
	str := L.CheckString(1)
	var data any
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		L.RaiseError("failed to decode JSON string: %v", err)
		return 0
	}
	L.Push(goValueToLua(L, data))
	return 1
}

// convert Go value to Lua value recursively
func goValueToLua(L *lua.LState, v any) lua.LValue {
	switch val := v.(type) {
	case string:
		return lua.LString(val)
	case float64:
		return lua.LNumber(val)
	case bool:
		return lua.LBool(val)
	case map[string]any:
		tbl := L.NewTable()
		for k, v2 := range val {
			tbl.RawSetString(k, goValueToLua(L, v2))
		}
		return tbl
	case []any:
		tbl := L.NewTable()
		for i, item := range val {
			// Lua tables index starting at 1
			tbl.RawSetInt(i+1, goValueToLua(L, item))
		}
		return tbl
	default:
		return lua.LNil
	}
}
