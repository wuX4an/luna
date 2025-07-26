package random

import (
	"math/rand"

	lua "github.com/yuin/gopher-lua"
)

func Choice(L *lua.LState) int {
	list := L.CheckTable(1)
	n := list.Len()

	if n == 0 {
		L.RaiseError("random.choice: table is empty")
		return 0
	}

	// Sin pesos: elección uniforme
	if L.GetTop() < 2 || L.Get(2).Type() != lua.LTTable {
		index := rand.Intn(n) + 1
		L.Push(list.RawGetInt(index))
		return 1
	}

	// Con pesos
	weights := L.CheckTable(2)
	if weights.Len() != n {
		L.RaiseError("random.choice: weights length must match list length")
		return 0
	}

	cumulative := make([]int, n)
	totalWeight := 0

	for i := 1; i <= n; i++ {
		wv := weights.RawGetInt(i)
		w, ok := wv.(lua.LNumber)
		if !ok {
			L.RaiseError("random.choice: weight at index %d is not a number", i)
			return 0
		}
		weight := int(w)
		if weight < 0 {
			L.RaiseError("random.choice: weight at index %d is negative", i)
			return 0
		}
		totalWeight += weight
		cumulative[i-1] = totalWeight
	}

	if totalWeight == 0 {
		L.RaiseError("random.choice: total weight is zero")
		return 0
	}

	r := rand.Intn(totalWeight) + 1
	for i, w := range cumulative {
		if r <= w {
			L.Push(list.RawGetInt(i + 1))
			return 1
		}
	}

	L.RaiseError("random.choice: failed to select value")
	return 0
}
