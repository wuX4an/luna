// map.go
package tablex

import (
	lua "github.com/yuin/gopher-lua"
)

// tablex.map(tbl, fn)
func Map(L *lua.LState) int {
	tbl := L.CheckTable(1)
	fn := L.CheckFunction(2)
	newTbl := L.NewTable()

	tbl.ForEach(func(k, v lua.LValue) {
		L.Push(fn)
		L.Push(v)
		L.Push(k)
		if err := L.PCall(2, 1, nil); err != nil {
			L.RaiseError("tablex.map error: %v", err)
		}
		newTbl.RawSet(k, L.Get(-1))
		L.Pop(1)
	})

	L.Push(newTbl)
	return 1
}
