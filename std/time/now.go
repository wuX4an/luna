package time

import (
	"time"

	lua "github.com/yuin/gopher-lua"
)

func now(L *lua.LState) int {
	L.Push(lua.LNumber(time.Now().Unix()))
	return 1
}
