package json

import (
	"encoding/json"

	lua "github.com/yuin/gopher-lua"
)

// Encode: converts Lua table to JSON string
func Encode(L *lua.LState) int {
	tbl := L.CheckTable(1) // check first argument is a table

	// convert Lua table to Go map[string]interface{}
	goMap := make(map[string]any)
	tbl.ForEach(func(k, v lua.LValue) {
		goMap[k.String()] = luaValueToGo(v)
	})

	b, err := json.Marshal(goMap)
	if err != nil {
		L.RaiseError("failed to encode table to JSON: %v", err)
		return 0
	}

	L.Push(lua.LString(b))
	return 1
}

// helper: convert Lua value to Go value
func luaValueToGo(v lua.LValue) any {
	switch v.Type() {
	case lua.LTBool:
		return lua.LVAsBool(v)
	case lua.LTNumber:
		return float64(v.(lua.LNumber))
	case lua.LTString:
		return v.String()
	case lua.LTTable:
		tbl := v.(*lua.LTable)
		m := make(map[string]any)
		tbl.ForEach(func(k, v lua.LValue) {
			m[k.String()] = luaValueToGo(v)
		})
		return m
	default:
		return nil
	}
}
