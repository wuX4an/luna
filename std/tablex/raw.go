package tablex

import (
	"strings"

	lua "github.com/yuin/gopher-lua"
)

func Raw(L *lua.LState) int {
	tbl := L.CheckTable(1)
	visited := map[*lua.LTable]bool{}
	str := dumpTable(tbl, visited)
	L.Push(lua.LString(str))
	return 1
}

func dumpTable(tbl *lua.LTable, visited map[*lua.LTable]bool) string {
	if visited[tbl] {
		return "<recursive>"
	}
	visited[tbl] = true

	var items []string

	tbl.ForEach(func(k, v lua.LValue) {
		var item strings.Builder
		item.WriteString(toString(k))
		item.WriteString(" = ")
		switch lv := v.(type) {
		case *lua.LTable:
			item.WriteString(dumpTable(lv, visited))
		default:
			item.WriteString(toString(lv))
		}
		items = append(items, item.String())
	})

	return "{ " + strings.Join(items, ", ") + " }"
}

func toString(val lua.LValue) string {
	if val.Type() == lua.LTString {
		return `"` + val.String() + `"`
	}
	return val.String()
}
