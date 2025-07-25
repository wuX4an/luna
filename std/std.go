package std

import (
	"luna/std/time"

	lua "github.com/yuin/gopher-lua"
)

func RegisterAll(L *lua.LState) {
	L.PreloadModule("time", time.Loader)
}
