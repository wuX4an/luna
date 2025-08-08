package math

import (
	"math"

	lua "github.com/yuin/gopher-lua"
)

func Pow(L *lua.LState) int {
	nargs := L.GetTop()
	if nargs != 2 {
		L.RaiseError(
			"too many arguments: expected exactly 2 arguments (base, exponent), got %d",
			nargs,
		)

		return 0
	}

	base := L.CheckNumber(1)
	exponent := L.CheckNumber(2)

	power := math.Pow(float64(base), float64(exponent))
	L.Push(lua.LNumber(power))
	return 1
}
