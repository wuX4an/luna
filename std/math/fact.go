package math

import (
	lua "github.com/yuin/gopher-lua"
)

// Fact calcula el factorial de un entero no negativo.
// Se expone a Lua como math.fact.
func Fact(L *lua.LState) int {
	nargs := L.GetTop()

	if nargs != 1 {
		L.RaiseError("expected exactly 1 argument, got %d", nargs)
		return 0
	}

	val := L.CheckNumber(1)
	n := int(val)
	if float64(n) != float64(val) || n < 0 {
		L.RaiseError("argument must be a non-negative integer, got %v", val)
		return 0
	}

	result := factorial(n)
	L.Push(lua.LNumber(result))
	return 1
}

func factorial(n int) float64 {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1.0
	for i := 2; i <= n; i++ {
		result *= float64(i)
	}
	return result
}
