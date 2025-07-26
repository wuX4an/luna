// filter.go
package tablex

import (
	lua "github.com/yuin/gopher-lua"
)

// tablex.filter(tbl, fn)
func Filter(L *lua.LState) int {
	tbl := L.CheckTable(1)
	fn := L.CheckFunction(2)
	newTbl := L.NewTable()

	tbl.ForEach(func(k, v lua.LValue) {
		L.Push(fn)
		L.Push(v)
		L.Push(k)
		if err := L.PCall(2, 1, nil); err != nil {
			L.RaiseError("tablex.filter error: %v", err)
		}
		ret := L.Get(-1)
		L.Pop(1)

		if lua.LVAsBool(ret) {
			newTbl.RawSet(k, v)
		}
	})

	L.Push(newTbl)
	return 1
}
