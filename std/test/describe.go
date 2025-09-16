//go:build !js

package test

import (
	"fmt"
	"os"

	lua "github.com/yuin/gopher-lua"
)

// test.describe("nombre", function() ... end)
func Describe(L *lua.LState) int {
	desc := L.CheckString(1)
	fn := L.CheckFunction(2)

	// Si LUNA_RUN_TESTS != "1", solo registrar tests, no ejecutar
	if os.Getenv("LUNA_RUN_TESTS") != "1" {
		tbl := L.NewTable()
		L.SetField(tbl, "type", lua.LString("describe"))
		L.SetField(tbl, "desc", lua.LString(desc))
		L.SetField(tbl, "fn", fn)

		return 0
	}

	// Ejecutar describe inmediatamente
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
