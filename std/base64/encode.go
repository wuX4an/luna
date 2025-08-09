package base64

import (
	"encoding/base64"

	lua "github.com/yuin/gopher-lua"
)

func Encode(L *lua.LState) int {
	nargs := L.GetTop()
	if nargs != 1 {
		L.RaiseError(
			"too many arguments: expected exactly 1 arguments (string), got %d",
			nargs,
		)
		return 0
	}

	input := L.CheckString(1)
	encoded := base64.URLEncoding.EncodeToString([]byte(input))

	L.Push(lua.LString(encoded))
	return 1
}
