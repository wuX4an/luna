package std

import (
	"luna/std/base64"
	"luna/std/crypto"
	"luna/std/env"
	"luna/std/http"
	"luna/std/ipc"
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
	L.PreloadModule("std:ipc", ipc.Loader)
	L.PreloadModule("std:base64", base64.Loader)
	L.PreloadModule("std:env", env.Loader)
	L.PreloadModule("std:crypto", crypto.Loader)
}
