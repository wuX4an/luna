package std

import (
	"luna/std/http"
	"luna/std/math"
	"luna/std/random"
	"luna/std/sqlite"
	"luna/std/tablex"
	"luna/std/test"
	"luna/std/time"

	lua "github.com/yuin/gopher-lua"
)

func RegisterAll(L *lua.LState) {
	L.PreloadModule("std:time", time.Loader)
	L.PreloadModule("std:random", random.Loader)
	L.PreloadModule("std:tablex", tablex.Loader)
	L.PreloadModule("std:sqlite", sqlite.Loader)
	L.PreloadModule("std:test", test.Loader)
	L.PreloadModule("std:http", http.Loader)
	L.PreloadModule("std:math", math.Loader)
}
