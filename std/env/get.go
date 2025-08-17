package env

import (
	"os"

	lua "github.com/yuin/gopher-lua"
)

func Get(L *lua.LState) int {
	nargs := L.GetTop()
	if nargs != 1 {
		L.RaiseError(
			"expected exactly 1 argument (string), got %d",
			nargs,
		)
		return 0
	}

	key := L.CheckString(1)
	value, exists := os.LookupEnv(key)
	if !exists {
		L.Push(lua.LNil)
		return 1
	}

	L.Push(lua.LString(value))
	return 1
}
