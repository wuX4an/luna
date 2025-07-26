package time

import (
	"time"

	lua "github.com/yuin/gopher-lua"
)

func Sleep(L *lua.LState) int {
	ms := L.CheckInt(1)
	time.Sleep(time.Duration(ms) * time.Millisecond)
	return 0
}
