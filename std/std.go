package std

import (
	"luna/std/random"
	"luna/std/tablex"
	"luna/std/time"

	lua "github.com/yuin/gopher-lua"
)

func RegisterAll(L *lua.LState) {
	L.PreloadModule("time", time.Loader)
	L.PreloadModule("random", random.Loader)
	L.PreloadModule("tablex", tablex.Loader)
}
