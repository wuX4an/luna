package random

import (
	"math/rand"

	lua "github.com/yuin/gopher-lua"
)

func Int(L *lua.LState) int {
	nargs := L.GetTop()

	// Requiere al menos min y max
	if nargs < 2 {
		L.ArgError(1, "expected at least 2 arguments: min, max")
		return 0
	}

	min := L.CheckInt(1)
	max := L.CheckInt(2)

	if min > max {
		L.ArgError(1, "min must be <= max")
		return 0
	}

	skip := 1
	if nargs >= 3 {
		s := L.CheckInt(3)
		if s <= 0 {
			L.ArgError(3, "skip must be > 0")
			return 0
		}
		skip = s
	}

	// Generar todos los valores posibles
	var values []int
	for i := min; i <= max; i += skip {
		values = append(values, i)
	}

	if len(values) == 0 {
		L.ArgError(1, "no valid values in range")
		return 0
	}

	// Seleccionar uno aleatoriamente
	selected := values[rand.Intn(len(values))]
	L.Push(lua.LNumber(selected))
	return 1
}
