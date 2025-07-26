package random

import (
	"math"
	"math/rand"

	lua "github.com/yuin/gopher-lua"
)

func Float(L *lua.LState) int {
	nargs := L.GetTop()

	if nargs < 2 {
		L.ArgError(1, "expected at least 2 arguments: min, max")
		return 0
	}

	min := L.CheckNumber(1)
	max := L.CheckNumber(2)

	if min > max {
		L.ArgError(1, "min must be <= max")
		return 0
	}

	decimals := 2 // default max_digits
	if nargs >= 3 {
		decimals = L.CheckInt(3)
		if decimals < 0 {
			L.ArgError(3, "max_digits must be >= 0")
			return 0
		}
	}

	scale := math.Pow(10, float64(decimals))
	r := min + lua.LNumber(rand.Float64())*(max-min)
	rounded := math.Round(float64(r)*scale) / scale

	L.Push(lua.LNumber(rounded))
	return 1
}
