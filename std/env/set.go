//go:build !js

package env

import (
	"os"

	lua "github.com/yuin/gopher-lua"
)

func Set(L *lua.LState) int {
	nargs := L.GetTop()
	if nargs != 2 {
		L.RaiseError("expected exactly 2 arguments: key (string), value (string), got %d", nargs)
		return 0
	}

	key := L.CheckString(1)
	value := L.CheckString(2)

	vars[key] = value
	os.Setenv(key, value)

	L.Push(lua.LTrue)
	return 1
}
