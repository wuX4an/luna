package math

import (
	"math"

	lua "github.com/yuin/gopher-lua"
)

func Sin(L *lua.LState) int {
	nargs := L.GetTop()

	// Validar número de argumentos
	if nargs != 2 {
		L.RaiseError(
			"expected exactly 2 arguments: number, unit ('r' or 'g'); got %d",
			nargs,
		)
		return 0
	}

	// Extraer argumentos
	lnumber := L.CheckNumber(1)
	lstring := L.CheckString(2)

	// Validar unidad
	if lstring != "r" && lstring != "g" {
		L.RaiseError(
			"invalid unit: expected 'r' for radians or 'g' for degrees; got '%s'",
			lstring,
		)
		return 0
	}

	var result float64

	// Calcular según unidad
	switch lstring {
	case "r":
		result = math.Sin(float64(lnumber))
	case "g":
		result = math.Sin(float64(lnumber) * math.Pi / 180)
	}

	L.Push(lua.LNumber(result))
	return 1
}
