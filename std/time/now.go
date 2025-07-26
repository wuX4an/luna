package time

import (
	"time"

	lua "github.com/yuin/gopher-lua"
)

func Now(L *lua.LState) int {
	L.Push(lua.LNumber(time.Now().Unix()))
	return 1
}
