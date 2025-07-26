package random

import (
	"math/rand"

	lua "github.com/yuin/gopher-lua"
)

// Take recibe una tabla y un número N, devuelve N elementos aleatorios sin repetición.
func Take(L *lua.LState) int {
	list := L.CheckTable(1)
	n := list.Len()
	count := L.CheckInt(2)

	if count < 0 {
		L.ArgError(2, "must be non-negative")
		return 0
	}
	if count > n {
		L.ArgError(2, "cannot take more elements than present in list")
		return 0
	}

	// Crea slice con índices 1..n
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i + 1
	}

	// Barajear índices (Fisher-Yates)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		indices[i], indices[j] = indices[j], indices[i]
	}

	// Crea nueva tabla Lua con los elementos seleccionados
	result := L.CreateTable(count, 0)
	for i := 0; i < count; i++ {
		val := list.RawGetInt(indices[i])
		result.RawSetInt(i+1, val)
	}

	L.Push(result)
	return 1
}
