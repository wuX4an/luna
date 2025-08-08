package math

import (
	"math"

	lua "github.com/yuin/gopher-lua"
)

func Sqrt(L *lua.LState) int {
	nargs := L.GetTop()

	if nargs != 1 {
		L.RaiseError("expected exactly 1 argument, got %d", nargs)
		return 0
	}

	val := L.CheckNumber(1)
	if val < 0 {
		L.RaiseError("math domain error: sqrt of negative number %v is undefined", float64(val))
		return 0
	}

	result := math.Sqrt(float64(val))
	L.Push(lua.LNumber(result))
	return 1
}
