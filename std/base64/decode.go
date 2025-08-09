package base64

import (
	"encoding/base64"

	lua "github.com/yuin/gopher-lua"
)

func Decode(L *lua.LState) int {
	nargs := L.GetTop()
	if nargs != 1 {
		L.RaiseError(
			"too many arguments: expected exactly 1 argument (string), got %d",
			nargs,
		)
		return 0
	}

	input := L.CheckString(1)

	decoded, err := base64.URLEncoding.DecodeString(input)
	if err != nil {
		L.RaiseError("failed to decode base64 string: %v", err)
		return 0
	}

	L.Push(lua.LString(decoded))
	return 1
}
