package test

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

// test.describe("nombre", function() ... end)
func Describe(L *lua.LState) int {
	desc := L.CheckString(1)
	fn := L.CheckFunction(2)

	// Imprime el nombre del grupo de pruebas
	fmt.Printf("\n\033[33m[%s]\033[0m\n", desc)

	err := L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    0,
		Protect: true,
	})

	if err != nil {
		fmt.Printf("Error in describe block: %v\n", err)
	}

	return 0
}
